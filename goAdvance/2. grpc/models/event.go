package models

type Event struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date"`
	Location    string `json:"location"`
}

var Events = []Event{
    {ID: 1, Title: "GoLang Workshop", Description: "Learn Go from scratch", Date: "2024-01-10", Location: "Jakarta"},
    {ID: 2, Title: "Tech Conference 2024", Description: "Discuss the latest in tech", Date: "2024-02-15", Location: "Surabaya"},
    {ID: 3, Title: "Startup Meetup", Description: "Connect with startups and VCs", Date: "2024-03-05", Location: "Bandung"},
    {ID: 4, Title: "AI Summit", Description: "Deep dive into Artificial Intelligence", Date: "2024-04-20", Location: "Yogyakarta"},
    {ID: 5, Title: "Cloud Computing Expo", Description: "Explore the world of cloud tech", Date: "2024-05-18", Location: "Jakarta"},
    {ID: 6, Title: "Blockchain Seminar", Description: "Understand blockchain technology", Date: "2024-06-25", Location: "Bali"},
    {ID: 7, Title: "Cybersecurity Forum", Description: "Best practices in cybersecurity", Date: "2024-07-12", Location: "Surabaya"},
    {ID: 8, Title: "Mobile Dev Workshop", Description: "Hands-on mobile app development", Date: "2024-08-22", Location: "Jakarta"},
}