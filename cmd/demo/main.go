package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/riclib/open-props-css/uicss"
)

func main() {
	// Serve the CSS file
	http.Handle("/static/dashboard.css", uicss.CSSHandler())
	
	// Serve the interactive stylebook
	http.Handle("/stylebook", uicss.StylebookHandler())
	
	// Serve a demo dashboard
	http.HandleFunc("/", dashboardHandler)
	
	fmt.Println("Demo server running at http://localhost:8080")
	fmt.Println("- Dashboard: http://localhost:8080/")
	fmt.Println("- Stylebook: http://localhost:8080/stylebook")
	
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Dashboard Demo</title>
    <link rel="stylesheet" href="/static/dashboard.css">
</head>
<body>
    <nav>
        <h1>My Dashboard</h1>
        <ul>
            <li><a href="#" class="active">Overview</a></li>
            <li><a href="#">Analytics</a></li>
            <li><a href="#">Reports</a></li>
            <li><a href="#">Settings</a></li>
        </ul>
    </nav>
    
    <main>
        <header>
            <h2>Overview</h2>
            <button class="primary">New Item</button>
        </header>
        
        <div class="grid grid-cols-3 gap-4 mb-4">
            <div class="card">
                <h4>Total Users</h4>
                <p class="text-large font-bold">1,234</p>
                <p class="text-muted text-small">+12% from last month</p>
            </div>
            <div class="card">
                <h4>Revenue</h4>
                <p class="text-large font-bold">$45,678</p>
                <p class="text-muted text-small">+8% from last month</p>
            </div>
            <div class="card">
                <h4>Active Projects</h4>
                <p class="text-large font-bold">56</p>
                <p class="text-muted text-small">3 completed this week</p>
            </div>
        </div>
        
        <div class="card">
            <h3>Recent Activity</h3>
            <table>
                <thead>
                    <tr>
                        <th>User</th>
                        <th>Action</th>
                        <th>Date</th>
                        <th>Status</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                        <td>John Doe</td>
                        <td>Created new project</td>
                        <td>2 hours ago</td>
                        <td><span class="badge success">Completed</span></td>
                    </tr>
                    <tr>
                        <td>Jane Smith</td>
                        <td>Updated settings</td>
                        <td>5 hours ago</td>
                        <td><span class="badge primary">In Progress</span></td>
                    </tr>
                    <tr>
                        <td>Bob Johnson</td>
                        <td>Submitted report</td>
                        <td>1 day ago</td>
                        <td><span class="badge success">Completed</span></td>
                    </tr>
                    <tr>
                        <td>Alice Brown</td>
                        <td>Requested access</td>
                        <td>2 days ago</td>
                        <td><span class="badge warning">Pending</span></td>
                    </tr>
                </tbody>
            </table>
        </div>
        
        <div class="flex gap-2 justify-end mt-4">
            <button>Previous</button>
            <button class="primary">Next</button>
        </div>
    </main>
</body>
</html>`
	
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}