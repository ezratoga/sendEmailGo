### Send Email And Attachment
This is web page and backend system for sending email and attachment.

### Features
You can input you Subject, Email Content and Email Attachment

### Tech Stack
- Go (Golang): Core backend language
- Echo: Web framework for routing and middleware
- HTML/CSS: CSS for styling and HTML for displaying the content of web page
- Froala Text Editor: For editing and styling text into your email and attachment. You must be connected to the internet if you want to send email by webpage, so the Froala form can be shown in the web page
- Javascript: To make web page more responsive and dynamic

### Environment Variables
This is the key that yo can declare in .env file:

Variable = Description

- SMTP_EMAIL = Your SMTP email address
- SMTP_PASSWORD = YOUR SMTP Password
- SMTP_HOST = Your SMTP Domain
- SMTP_PORT = Your SMTP PORT
- EMAIL_RECEIVER = The target of email recieiver

### Setup and Installation

Clone the repository:
```shell
$ git clone https://github.com/ezratoga/sendEmailGo.git
$ cd dating-app-backend
```

Install dependencies:
```shell
$ go mod download
```

Set up the environment variables: Create a .env file in the root directory and add the required environment variables as listed above. The port is hardcoded run at `8899`

Run the application in your local cmd:
```shell
$ go run main.go
```

### To Access the webpage
After you run the backend service, you can open file `wysiwig.html` in this repository to accessing web page

