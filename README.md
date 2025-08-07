# REST Blog API

This is a simple RESTful API for managing personal blog articles.
Built with Go & SQL, supports full CRUD functionality with basic filtering: by title, start date, & end date.

---

## Features

- CREATE new article
- FILTER articles by:
  - Title
  - Start date
  - End date
  *If you didn't set any filters here, it will retrieved all articles.
- FIND article by ID
- UPDATE an article (title & content) by ID
- DELETE an article by ID

---

## Tech Stack

- **Language**: Go
- **Database**: MySQL
- **Driver**: `go-sql-driver/mysql`
- **Environment Config**: `.env` with `github.com/joho/godotenv`

---

## Folder Structure

```
rest-blog-api/
├── app/            # Database connection & config
├── controller/     # HTTP request handlers
├── exception/      # Error handler
├── helper/         # Common utilities (PanicIfError, JSON Decode/Encode, etc)
├── model/          # Domain & web models (Struct)
├── repository/     # DB query logic
├── service/        # Business logic
├── main.go         # Main func (router, dependency injection)
├── openapi.json    # API spec documentation
└── README.md       # Project info
```

---

## Getting Started

### 1. Clone the repo
```bash
git clone https://github.com/pisondev/rest-blog-api.git
cd rest-blog-api
```

### 2. Create a `.env` file

```env
DB_USER=root
DB_PASS=yourpassword
DB_HOST=localhost
DB_PORT=3306
DB_NAME=rest_blog_api
DB_PARAMS=parseTime=true&loc=Asia%2FJakarta
```

> ⚠️ Don't commit your `.env` file — it's already in `.gitignore`.
If you're using an older commit and hit DB errors, copy the latest app/ folder — it has updated .env-based config.

### 3. Install dependencies
```bash
go mod tidy
```

### 4. Run the app
```bash
go run main.go
```

---

## API Endpoints

| Method | Endpoint           | Description                     |
|--------|--------------------|---------------------------------|
| POST   | `/articles`        | Create a new article            |
| GET    | `/articles`        | Get all articles (with filters) |
| GET    | `/articles/:id`    | Get article by ID               |
| PUT    | `/articles/:id`    | Update article by ID            |
| DELETE | `/articles/:id`    | Delete article by ID            |

> Filtering supported via query params: `title`, `start_date`, `end_date`

---

## TODO / Roadmap

- [ ] Add user authentication (JWT)
- [ ] Assign articles to users
- [ ] Add tags/categories support
- [ ] Pagination and sorting
- [ ] Tests and CI/CD
- [ ] Deployment

---

## License

MIT — free to use and modify.
