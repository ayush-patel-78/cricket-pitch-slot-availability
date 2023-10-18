
package main

import(
    "net/http"
    "fmt"
    "io/ioutil"
	"time"
	"net/url"
	"encoding/json"
)


// entities

type Resource struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type BusinessHour struct {
	Id         string    `json:"id"`
	ResourceId string    `json:"resource_id"`
	Quantity   int64     `json:"quantity"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}

type BlockHour struct {
	Id         string    `json:"id"`
	ResourceId string    `json:"resource_id"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}

type Appointment struct {
	Id         string    `json:"id"`
	ResourceId string    `json:"resource_id"`
	Quantity   int64     `json:"quantity"`
	StartTime  string    `json:"start_time"`
	EndTime    string    `json:"end_time"`
}

type Duration struct {
	Seconds int64 `json:"seconds"`
}

// Data

type AvailableBusinessHour struct {
	StartTime time.Time
	EndTime   time.Time
	Quantity  int64
}

type AvailableSlots struct {
	StartTime time.Time
	EndTime   time.Time
}
// endpoint request structs

type ListBusinessHoursRequest struct {
	ResourceId string `json:"resourceId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

type ListBlockHoursRequest struct {
	ResourceId string `json:"resourceId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

type ListAppointmentRequest struct {
	ResourceId string `json:"resourceId"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
}

// helper functions

func TimeToString(tm time.Time) string {
	return tm.Format(time.RFC3339)
}

func StringToTime(timeStr string) time.Time {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		return time.Time{}
	}

	return t
}


// main
// make an api call to get the list of durations
// make an api call to get the list of business hours
// make an api call to get the list of block hours
// make an api call to get the list of appointments
// find the available slots
// return the available slots

// getDurationList function will give the list of durations

func getDurationList() []Duration{
	url := "http://api.internship.appointy.com:8000/v1/durations"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
			fmt.Print(err.Error())
	}
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA4LTEwVDAwOjAwOjAwWiIsInVzZXJfaWQiOjE5N30.mUSopGe9mTOE_YcCMD8Bev0pCT1zxA1zwTmSzkkjOsA")
	res, err := http.DefaultClient.Do(req)
	// fmt.Println(res)
	if err != nil {
			fmt.Print(err.Error())
			return nil
	}
	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
			fmt.Print(err.Error())
			return nil
	}
	// fmt.Println(string(body))

	// unmarshal the json data
	var durations []Duration
	json.Unmarshal([]byte(body), &durations)
	// for _, duration := range durations {
	// 	fmt.Println(duration.Seconds)
	// }
    return durations

}

// getbusinesshours function will give the list of business hours
   
func getbusinesshours(id string, start string, end string) []BusinessHour {

// passing parameters
	baseURL, _ := url.Parse("http://api.internship.appointy.com:8000/v1/business-hours")
    params := url.Values{}
    params.Add("resourceId", id)
    params.Add("startTime", start)
	params.Add("endTime", end)
    baseURL.RawQuery = params.Encode()
	req,err := http.NewRequest("GET", baseURL.String(), nil)
	if err != nil {
			fmt.Print(err.Error())
			return nil
	}
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA4LTEwVDAwOjAwOjAwWiIsInVzZXJfaWQiOjE5N30.mUSopGe9mTOE_YcCMD8Bev0pCT1zxA1zwTmSzkkjOsA")
    res, err := http.DefaultClient.Do(req)
	// fmt.Println(res)
	if err != nil {
			fmt.Print(err.Error())
	}
	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
			fmt.Print(err.Error())
	}
	// fmt.Println(string(body))

	// unmarshal the json data
	var businesshours []BusinessHour
	json.Unmarshal([]byte(body), &businesshours)

    return businesshours

}

// getappointments function will give the list of appointments

func getappointments(id string, start string, end string) []Appointment{

	// passing parameters
		baseURL, _ := url.Parse("http://api.internship.appointy.com:8000/v1/appointments")
		params := url.Values{}
		params.Add("resourceId", id)
		params.Add("startTime", start)
		params.Add("endTime", end)
		baseURL.RawQuery = params.Encode()
		req,err := http.NewRequest("GET", baseURL.String(), nil)
		if err != nil {
				fmt.Print(err.Error())
		}
		req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA4LTEwVDAwOjAwOjAwWiIsInVzZXJfaWQiOjE5N30.mUSopGe9mTOE_YcCMD8Bev0pCT1zxA1zwTmSzkkjOsA")
		res, err := http.DefaultClient.Do(req)
		// fmt.Println(res)
		if err != nil {
				fmt.Print(err.Error())
		}
		defer res.Body.Close()
		body, readErr := ioutil.ReadAll(res.Body)
		if readErr != nil {
				fmt.Print(err.Error())
		}
		// fmt.Println(string(body))
		// unmarshal the json data
	var appointmentschedule []Appointment
	json.Unmarshal([]byte(body), &appointmentschedule)

	// for _, appointment := range appointmentschedule {
	// 	fmt.Println(appointment.StartTime)
	// 	fmt.Println(appointment.EndTime)
	// 	fmt.Println(appointment.Quantity)
	// 	fmt.Println(appointment.ResourceId)
	// 	fmt.Println(appointment.Id)
	// }

    return appointmentschedule
	
	}

// getresources function will give the list of resources
	
func getresources() []Resource{
	url := "http://api.internship.appointy.com:8000/v1/resources"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
			fmt.Print(err.Error())
	}
	req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA4LTEwVDAwOjAwOjAwWiIsInVzZXJfaWQiOjE5N30.mUSopGe9mTOE_YcCMD8Bev0pCT1zxA1zwTmSzkkjOsA")
	res, err := http.DefaultClient.Do(req)
	// fmt.Println(res)
	if err != nil {
			fmt.Print(err.Error())
	}
	defer res.Body.Close()
	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
			fmt.Print(err.Error())
	}
	// fmt.Println(string(body))
	// unmarshal the json data
	var resources []Resource
	json.Unmarshal([]byte(body), &resources)
	// for _, resource := range resources {
	// 	fmt.Println(resource.Id)
	// }

	
	return resources

}

// getblockhours function will give the list of block hours

func getblockhours(id string, start string, end string) []BlockHour{
// passing parameters
baseURL, _ := url.Parse("http://api.internship.appointy.com:8000/v1/block-hours")
params := url.Values{}
params.Add("resourceId", id)
params.Add("startTime", start)
params.Add("endTime", end)
baseURL.RawQuery = params.Encode()
req,err := http.NewRequest("GET", baseURL.String(), nil)
if err != nil {
		fmt.Print(err.Error())
}
req.Header.Add("Authorization", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOiIyMDIzLTA4LTEwVDAwOjAwOjAwWiIsInVzZXJfaWQiOjE5N30.mUSopGe9mTOE_YcCMD8Bev0pCT1zxA1zwTmSzkkjOsA")
res, err := http.DefaultClient.Do(req)
// fmt.Println(res)
if err != nil {
		fmt.Print(err.Error())
}
defer res.Body.Close()
body, readErr := ioutil.ReadAll(res.Body)
if readErr != nil {
		fmt.Print(err.Error())
}
// fmt.Println(string(body))
// unmarshal the json data
    var blockhours []BlockHour
	json.Unmarshal([]byte(body), &blockhours)
	
	return blockhours

}


func mapping(mpp map[string]int64,startTime string,endTime string,quantity int64){
	var temp time.Time = StringToTime(startTime)
	var end time.Time = StringToTime(endTime)
	// fmt.Println("interval of 30 mins is created by mapping function ")
	for ;temp.Before(end);temp = temp.Add(time.Duration(30) * time.Minute){
		// fmt.Println(TimeToString(temp))
		mpp[TimeToString(temp)] = quantity
	}
	
}

func getAvailability(id string, date string, time_required int64,quantity int64){
	time_required = time_required * 60
   var dur []Duration = getDurationList()
   var flag bool 
   flag = false
   for _, duration := range dur {
		if duration.Seconds == time_required {
			flag = true
			fmt.Println("time_required is valid")
		}
	}
	if flag == false {
		fmt.Println("Invalid time_required")
		return
	}

	var resources []Resource = getresources()
	var flag_res bool
	flag_res = false
	for _, resource := range resources {
		if resource.Id == id {
			flag_res = true
			fmt.Println("pitch is found")
		}
	}
	if flag_res == false {
		fmt.Println("ResourceId or pitch is invalid ");
		return 
	}

	var start_time string = date + "T00:00:00Z"
	var end_time string = date + "T23:59:59Z"

	// Appointment map work started from here

	var appointments []Appointment = getappointments(id,start_time,end_time)
	// i have to create a map of appointment key is string (which is startTime + 30 minutes) and value is quantity
	var appointmentmap = make(map[string]int64)
	
	
	for _, appointment := range appointments {
		// call mapping function
        mapping(appointmentmap,appointment.StartTime,appointment.EndTime,appointment.Quantity)

	}

	// iterate over appointment map 
	// fmt.Println("appointment map iteration:",appointmentmap)

	// Appointment map work completed 

	// Businesshours map work started from here 

	var businesshours []BusinessHour = getbusinesshours(id,start_time,end_time)
    var businesshourmap = make(map[string]int64)

	for _, business := range businesshours{
		// call mapping function 
		mapping(businesshourmap,business.StartTime,business.EndTime,business.Quantity)
	}

	// Iterate over businesshour map
	// fmt.Println("businesshour map iteration:",businesshourmap)

	// Business map work ended here


	// Blockhour map work started here

	var blockhours []BlockHour = getblockhours(id,start_time,end_time)
	var blockhourmap = make(map[string]int64)

	for _,block := range blockhours{
		// call mapping function
		mapping(blockhourmap,block.StartTime,block.EndTime,0)
	}

	// Iterate over blockhour map
	// fmt.Println("blockhour map iteration:", blockhourmap)
	// Blockhour map work completed




   // logic to find available slots starts from here 
   
   var checkslot int64 = time_required/1800

   var availableslotsmap = make(map[string]int64)

   // we have to iterate over business hour map and avoid slot which is present in the block hours 
   var count int64 = 0
  
   for businesshourStartTime,businesshourQuantity := range businesshourmap{
	  	// find businesshourStartTime should not collide with blockhour time 
		//   fmt.Println("hehe:",businesshourStartTime,businesshourQuantity)
	   	_,present := blockhourmap[businesshourStartTime]
			// fmt.Println("check block hour is present or not:",present)
			
			if present ==  false {
				seatsbooked,booked := appointmentmap[businesshourStartTime]
					if booked == true{
						var seatsLeft int64 = businesshourQuantity - seatsbooked
						if seatsLeft > quantity {
							count = count + 1
							if count > checkslot {
								var timeavailable time.Time = StringToTime(businesshourStartTime)
								timeavailable = timeavailable.Add(-time.Duration(30*(checkslot - 1)) * time.Minute)
								availableslotsmap[TimeToString(timeavailable)] = seatsLeft
							}

						} else {
							count = 0
						}
					} else{
						var seatsLeft int64 = businesshourQuantity
						if seatsLeft > quantity {
							count = count + 1
							if count > checkslot {
								var timeavailable time.Time = StringToTime(businesshourStartTime)
								timeavailable = timeavailable.Add(-time.Duration(30*(checkslot - 1)) * time.Minute)
								availableslotsmap[TimeToString(timeavailable)] = seatsLeft
							}

						} else {
							count = 0
						}

					}
				

			} else {
				count = 0
			}
		
   	}
   	fmt.Println("availableSlots map : ", availableslotsmap)
}





func main() {
	// getDurationList()
	// getbusinesshours("res_2","2021-08-04T00:00:00Z","2021-08-04T23:59:59Z")
	// getappointments("res_2","2021-08-04T00:00:00Z","2021-08-04T23:59:59Z")
	// getblockhours("res_2","2021-08-04T00:00:00Z","2021-08-04T23:59:59Z")
	// getresources()
    
	getAvailability("res_2","2023-08-04",60,1)

}


