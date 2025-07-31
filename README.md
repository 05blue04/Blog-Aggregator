# 🐊 Blog Aggregator — `gator`

A command-line blog aggregation tool written in Go that lets users follow RSS feeds, browse posts, and manage feed subscriptions directly from the terminal.

---

## 📦 Prerequisites

Make sure you have the following installed:

- **Go** `v1.24.5` or higher  
- **PostgreSQL** `v16.9` or higher  

---

## 🚀 Installation

Install the CLI tool using `go install`:

```bash
go install github.com/05blue04/Blog-Aggregator@latest
```
## 🛠️ Database Setup (Linux)

```bash
sudo service postgresql start
sudo -u postgres psql
```
Inside the psql shell:
```bash
CREATE DATABASE gator;
```

## ⚙️ Configuration
Create a .gatorconfig.json file in your home directory:

```bash
{
  "db_url": "postgres://postgres:postgres@localhost:5432/gator?sslmode=disable",
  "current_user_name": "your_username"
}
```
db_url — PostgreSQL connection string

current_user_name — your active gator user

## CLI Usage
| Command                     | Description                                  |
| --------------------------- | -------------------------------------------- |
| `gator reset`               | Resets the database (⚠️ deletes all data)    |
| `gator register "username"` | Registers a new user and sets them as active |
| `gator login "username"`    | Switches to an existing user                 |
| `gator users`               | Lists all registered users                   |
| `gator addfeed "url"`       | Adds a new RSS feed to your account          |
| `gator feeds`               | Displays all available feeds                 |
| `gator follow "url"`        | Follows a feed as the current user           |
| `gator following`           | Lists all feeds you're following             |
| `gator unfollow "url"`      | Unfollows a feed                             |
| `gator browse [limit]`       | Displays recent posts from followed feeds; optionally limit number of posts (default: 2) |

## Example
```bash
# Register and login
gator register alice
gator login alice

# Add a feed
gator addfeed https://xkcd.com/rss.xml

# Follow it
gator follow https://xkcd.com/rss.xml

# Browse posts
gator browse

```