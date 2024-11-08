# TrackPro

## Project Structure

TrackPro/
├── main.go # Entry point of the application
├── handlers/ # Contains HTTP handlers (controllers) for each API route
│ ├── project.go # Handlers related to project creation and summary
│ ├── requirement.go # Handlers for adding and listing requirements
│ └── effort.go # Handlers for logging and summarizing efforts
├── models/ # Contains the data models for Project, Requirement, and Effort
│ ├── project.go # Project model
│ ├── requirement.go # Requirement model
│ └── effort.go # Effort model
├── storage/ # In-memory data storage for all entities
│ └── storage.go # Basic CRUD functions for each model
├── utils/ # Utility functions, such as ID generation
│ └── utils.go
├── templates/ # HTML templates for rendering web pages
│ ├── index.html # Main page template
│ ├── project.html # Project details and summary template
│ └── form.html # Template for forms (e.g., create project, add requirement)
└── README.md # Project overview and setup instructions
