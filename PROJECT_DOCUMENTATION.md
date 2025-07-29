# 📦 Warehouse Management System - Project Documentation

## 🎯 Overview

This is a **Warehouse Management System** API built in **Go** that handles the complete logistics chain for product distribution. It manages warehouses, sellers, buyers, products, employees, and the entire order fulfillment process.

> **🎓 Context**: This is a bootcamp final project that simulates a real-world warehouse management system.

---

## 🏗️ System Architecture

### **📐 Clean Architecture Pattern**
The project follows **Clean Architecture** principles with clear separation of concerns:

```
ProyectoFinal/
├── 🎯 cmd/                     # Application entry point
├── 📊 docs/                    # Documentation & database
├── 🏢 internal/                # Core business logic
│   ├── 🚀 application/         # App configuration & setup
│   ├── 🌐 handler/             # HTTP handlers (controllers)
│   ├── 💾 repository/          # Data access layer
│   └── ⚙️  service/            # Business logic layer
├── 🧪 mocks/                   # Test mocks
└── 📦 pkg/                     # Shared packages
```

### **🔄 Request Flow**
```
HTTP Request → Handler → Service → Repository → Database
                ↓
HTTP Response ← Handler ← Service ← Repository ← Database
```

---

## 🏪 Business Domain

### **🎯 Core Entities**

| Entity | Description | Key Fields |
|--------|-------------|------------|
| **🏪 Seller** | Companies that sell products | CID, Company Name, Address, Locality |
| **🏭 Warehouse** | Storage facilities | Code, Address, Capacity, Temperature |
| **📦 Product** | Items being sold | Code, Description, Price, Category |
| **👥 Buyer** | Companies that purchase products | CID, Company Name, Purchase History |
| **👷 Employee** | Warehouse staff | Card Number, Name, Warehouse Assignment |
| **🚚 Carrier** | Transport companies | CID, Company Name, Routes |
| **📍 Locality** | Geographic locations | Name, Province, Country |
| **📋 Section** | Warehouse divisions | Number, Capacity, Temperature |

### **📊 Business Operations**

- **📥 Inbound Orders**: Products entering warehouses
- **🛒 Purchase Orders**: Buyer orders with multiple products
- **📦 Product Batches**: Grouped products by expiration/lot
- **📈 Product Records**: Tracking product movements
- **📊 Reports**: Locality analytics, carrier reports

---

## 🛠️ Technology Stack

### **🔧 Core Technologies**
```go
// Backend Framework
Go 1.24 + Chi Router + MySQL

// Key Dependencies
github.com/go-chi/chi/v5           // HTTP router
github.com/go-sql-driver/mysql     // MySQL driver  
github.com/go-playground/validator // Input validation
github.com/joho/godotenv           // Environment variables
```

### **🧪 Testing & Quality**
```go
github.com/stretchr/testify        // Testing framework
github.com/DATA-DOG/go-sqlmock     // Database mocking
```

---

## 🚀 Quick Start

### **📋 Prerequisites**
- Go 1.24+
- MySQL 8.0+
- Git

### **⚡ Installation**

```bash
# 1. Clone the repository
git clone <repository-url>
cd ProyectoFinal

# 2. Install dependencies
go mod tidy

# 3. Set up environment variables
cp .env.example .env
# Edit .env with your database credentials

# 4. Set up database
mysql -u root -p < docs/db/bootcamp_project.sql

# 5. Run the application
go run cmd/main.go
```

### **🌐 Server Access**
```
Server URL: http://localhost:8080
API Base:   http://localhost:8080/api/v1
```

---

## 📚 API Documentation

### **🔗 Available Endpoints**

| Module | Base Path | Description |
|--------|-----------|-------------|
| **🏪 Sellers** | `/api/v1/sellers` | Manage seller companies |
| **🏭 Warehouses** | `/api/v1/warehouses` | Warehouse operations |
| **📦 Products** | `/api/v1/products` | Product catalog |
| **👥 Buyers** | `/api/v1/buyers` | Buyer management |
| **👷 Employees** | `/api/v1/employees` | Staff management |
| **🚚 Carriers** | `/api/v1/carriers` | Transport companies |
| **📍 Localities** | `/api/v1/localities` | Geographic data |
| **📋 Sections** | `/api/v1/sections` | Warehouse sections |
| **📦 Product Batches** | `/api/v1/productBatches` | Batch management |
| **📥 Inbound Orders** | `/api/v1/inboundOrders` | Incoming shipments |
| **🛒 Purchase Orders** | `/api/v1/purchaseOrders` | Customer orders |
| **📈 Product Records** | `/api/v1/productRecords` | Product tracking |

### **📖 Example API Calls**

#### **Create a Seller**
```bash
POST /api/v1/sellers
Content-Type: application/json

{
  "cid": "CID123",
  "company_name": "Fresh Foods Inc",
  "address": "123 Business St",
  "telephone": "555-0123",
  "locality_id": 1
}
```

#### **Get All Warehouses**
```bash
GET /api/v1/warehouses
```

#### **Get Seller by ID**
```bash
GET /api/v1/sellers/1
```

---

## 🗄️ Database Schema

### **🌍 Geographic Structure**
```sql
Countries → Provinces → Localities
    ↓
Sellers, Carriers, Warehouses
```

### **📦 Product Flow**
```sql
Products → Product Batches → Sections → Warehouses
    ↓
Inbound Orders, Purchase Orders, Product Records
```

### **👥 People & Companies**
```sql
Employees → Warehouses
Buyers → Purchase Orders
Sellers → Products
Carriers → Localities
```

---

## 📁 Project Structure Detailed

### **🚀 Application Layer** (`internal/application/`)
```
app.go              # Main server configuration
config/             # Database & environment setup
di/                 # Dependency injection
loader/             # Data loading utilities
router/             # Route definitions
```

### **🌐 Handler Layer** (`internal/handler/`)
```
- HTTP request/response handling
- Input validation
- Error handling
- Route parameter extraction

Example: seller_handler.go
├── Create()     # POST /sellers
├── GetAll()     # GET /sellers  
├── GetById()    # GET /sellers/:id
├── Update()     # PATCH/PUT /sellers/:id
└── Delete()     # DELETE /sellers/:id
```

### **⚙️ Service Layer** (`internal/service/`)
```
- Business logic implementation
- Data validation
- Business rule enforcement
- Cross-entity operations

Example: seller_default.go
├── Create()     # Validate and create seller
├── GetAll()     # Retrieve all sellers
├── GetById()    # Find seller by ID
├── Update()     # Update seller data
└── Delete()     # Remove seller
```

### **💾 Repository Layer** (`internal/repository/`)
```
- Database operations
- SQL query execution
- Data mapping
- Error handling

Example: seller_mysql.go
├── Create()     # INSERT operations
├── GetAll()     # SELECT all
├── GetById()    # SELECT by ID
├── Update()     # UPDATE operations
└── Delete()     # DELETE operations
```

### **📦 Models** (`pkg/models/`)
```
- Data structures
- Validation rules
- JSON serialization
- Domain entities

Example: seller.go
├── Seller struct           # Domain model
├── SellerDoc struct        # API documentation
├── CreateSellerRequest     # Input DTO
└── UpdateSellerRequest     # Update DTO
```

---

## 🧪 Testing Strategy

### **📊 Test Coverage: 71.5%**

### **🔍 Test Types**
- **Unit Tests**: Service & Repository layers
- **Integration Tests**: Handler layer
- **Mocking**: Database & external dependencies

### **🚀 Running Tests**
```bash
# Run all tests
make test

# Run with coverage report
make cover

# Run coverage without mocks
make cover-total

# Run specific module tests
go test ./internal/service/seller/ -v
```

### **📋 Test Structure**
```go
// Table-driven tests pattern
func TestSellerService_Create(t *testing.T) {
    tests := []struct {
        name       string
        input      models.Seller
        mockReturn models.Seller
        mockError  error
        assertFunc func(t *testing.T, result models.Seller, err error)
    }{
        // Test cases...
    }
    // Test execution...
}
```

---

## ⚙️ Configuration

### **🔧 Environment Variables**
```bash
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=bootcamp_project

# Server Configuration
SERVER_PORT=8080
SERVER_HOST=localhost
```

### **📋 Makefile Commands**
```bash
make test           # Run all tests
make cover          # Generate coverage report with HTML
make cover-total    # Generate coverage excluding mocks
```

---

## 🔄 Development Workflow

### **🌟 Adding a New Feature**

1. **📝 Define Model** (`pkg/models/`)
```go
type NewEntity struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}
```

2. **💾 Create Repository** (`internal/repository/`)
```go
type Repository interface {
    Create(entity NewEntity) (NewEntity, error)
    GetById(id int) (*NewEntity, error)
    // ... other methods
}
```

3. **⚙️ Implement Service** (`internal/service/`)
```go
func (s *Service) Create(entity NewEntity) (NewEntity, error) {
    // Business logic
    return s.repository.Create(entity)
}
```

4. **🌐 Create Handler** (`internal/handler/`)
```go
func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
    // Handle HTTP request
}
```

5. **🧪 Write Tests**
```go
func TestNewEntityService_Create(t *testing.T) {
    // Test implementation
}
```

6. **🔗 Add Routes** (`internal/application/router/`)
```go
r.Post("/", handler.Create)
r.Get("/{id}", handler.GetById)
```

---

## 🐛 Common Issues & Solutions

### **🔌 Database Connection Issues**
```bash
# Check MySQL service
sudo systemctl status mysql

# Verify credentials in .env
DB_USER=root
DB_PASSWORD=your_actual_password
```

### **📦 Dependency Issues**
```bash
# Clean and reinstall modules
go clean -modcache
go mod download
go mod tidy
```

### **🧪 Test Failures**
```bash
# Run tests with verbose output
go test ./... -v

# Run specific failing test
go test ./internal/service/seller/ -run TestSpecificTest -v
```

---

## 📈 Performance Considerations

### **🚀 Optimization Tips**
- Use database indexing on frequently queried fields
- Implement pagination for large datasets
- Use connection pooling for database
- Cache frequently accessed data
- Validate input early in the handler layer

### **📊 Monitoring**
- Check API response times
- Monitor database query performance
- Track memory usage
- Log error patterns

---

## 🤝 Contributing

### **📋 Code Standards**
- Follow Go conventions and best practices
- Write comprehensive tests (aim for >80% coverage)
- Use meaningful variable and function names
- Document public APIs
- Follow the existing architecture patterns

### **🔄 Pull Request Process**
1. Create feature branch from `main`
2. Implement feature with tests
3. Run `make test` and `make cover`
4. Update documentation if needed
5. Submit pull request with clear description

---

## 📚 Additional Resources

### **📖 Learning Materials**
- [Go Documentation](https://golang.org/doc/)
- [Chi Router](https://github.com/go-chi/chi)
- [Testify](https://github.com/stretchr/testify)
- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)

### **🛠️ Development Tools**
- [Postman](https://www.postman.com/) - API testing
- [MySQL Workbench](https://dev.mysql.com/downloads/workbench/) - Database management
- [Mockery](https://github.com/vektra/mockery) - Mock generation

---

## ❓ FAQ

**Q: How do I reset the database?**
```bash
mysql -u root -p < docs/db/bootcamp_project.sql
```

**Q: How do I add a new endpoint?**
Follow the "Adding a New Feature" section above.

**Q: Why are my tests failing?**
Check the "Common Issues & Solutions" section.

**Q: How do I check code coverage?**
```bash
make cover-total
```

---

**🎉 Happy Coding! This warehouse management system is designed to be scalable, maintainable, and educational.** 