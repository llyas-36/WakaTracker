<<<<<<< HEAD
// package main
=======
 package main
<<<<<<< HEAD:main.txt
import "fmt"
>>>>>>> 615a439 (Update main.go)
=======
>>>>>>> c3b4fb3 (Update and rename main.txt to main.go):main.go
 
// import (
//     "encoding/json"
//     "fmt"
//     "io"
//     "log"
//     "net/http"
// )

// func handle(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")
//     username := ""
//     password := ""
//     credentials := fmt.Sprintf("%s:%s", username, password)
//     // encodedCredentials := base64.StdEncoding.EncodeToString([]byte(credentials))

//     url := "https://wakatime.com/api/v1/users/current/stats/last_7_days"

<<<<<<< HEAD
//     req, err := http.NewRequest("GET", url, nil)
//     if err != nil {
//         log.Fatal(err)
//     }
//     req.Header.Set("Authorization", "Basic "+credentials)
=======
// 	url := ""
>>>>>>> 03ebf3b (handle internal)

//     client := http.Client{}
//     res, err := client.Do(req)
//     if err != nil {
//         log.Fatal(err)
//     }

//     body, err := io.ReadAll(res.Body)
//     if err != nil {
//         log.Fatal(err)
//     }

//     fmt.Fprintf(w, string(body))
// }

// func Day_7(w http.ResponseWriter, r *http.Request) {
//     w.Header().Set("Content-Type", "application/json")
//     url := "https://wakatime.com/api/v1/users/current/stats/last_7_days"
//     auth := fmt.Sprintf("%s:%s", "", "")

//     req, err := http.NewRequest("GET", url, nil)
//     if err != nil {
//         panic(err)
//     }

//     req.Header.Set("Authorization", "Basic "+auth)

//     client := http.DefaultClient
//     res, err := client.Do(req)
//     if err != nil {
//         panic(err)
//     }
//     body, err := io.ReadAll(res.Body)
//     if err != nil {
//         panic(err)
//     }
//     defer res.Body.Close()

//     var data map[string]any
//     err = json.Unmarshal(body, &data)
//     if err != nil {
//         panic(err)
//     }

//     assertData := data["data"].(map[string]interface{})
//     projects := assertData["projects"].([]interface{})
//     jsonDD, err := json.Marshal(projects)
//     if err != nil {
//         panic(err)
//     }

//     type Activity struct {
//         Name    string  `json:"name"`
//         Percent float64 `json:"percent"`
//         Text    string  `json:"text"`
//     }

//     var activity []Activity
//     jsonerr := json.Unmarshal([]byte(jsonDD), &activity)
//     if jsonerr != nil {
//         panic(jsonerr)
//     }

//     for _, a := range activity {
//         fmt.Printf("%s → %s\nPercent → %v\n", a.Name, a.Text, a.Percent)
//     }

//     container, _ := json.Marshal(activity)
//     fmt.Fprintf(w, string(container))
// }

// func dailyAverage(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
//     w.Header().Set("Content-Type", "application/json")
//     url := "https://wakatime.com/api/v1/users/current/stats/last_7_days"
//     auth := fmt.Sprintf("%s:%s", "", "")
=======
// 	w.Header().Set("Content-Type", "application/json")
<<<<<<< HEAD
// 	url := "https://wakatime.com/api/v1/users/current/stats/last_7_days"
// 	auth := fmt.Sprintf("%s:%s", ", "")
>>>>>>> 7c051b8 (update main.go)
=======
// 	url := ""
// 	
>>>>>>> e4f6d42 (modify internal)
>>>>>>> 03ebf3b (handle internal)

//     req, err := http.NewRequest("GET", url, nil)
//     if err != nil {
//         panic(err)
//     }

//     req.Header.Set("Authorization", "Basic "+auth)
// }

// func main() {
//     http.HandleFunc("/", handle)
//     fmt.Println("server running")
//     http.ListenAndServe(":8080", nil)
// }
