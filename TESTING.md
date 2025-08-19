# 🧪 Testing Guide - Locality & Seller

## 📋 Available Tests Summary

### **Locality Module**
- ✅ Handler Tests
- ✅ Service Tests
- ✅ Repository Tests

### **Seller Module**
- ✅ Handler Tests
- ✅ Service Tests
- ✅ Repository Tests

---

## 🚀 Execution Commands (Makefile)

### **🏃‍♂️ Run ALL project tests**
bash
make test

### **📊 Run with Coverage + HTML Report**
bash
make cover

### **📈 Total Coverage (excluding mocks)**
bash
make cover-total

---

## 🎯 Tests by Module

### **🗺️ LOCALITY TESTS**

#### **Handler Tests**
bash
go test ./internal/handler/locality/ -v

#### **Service Tests**
bash
go test ./internal/service/locality/ -v

#### **Repository Tests**
bash
go test ./internal/repository/locality/ -v

#### **All Locality tests**
bash
go test ./internal/handler/locality/ ./internal/service/locality/ ./internal/repository/locality/ -v

---

### **👤 SELLER TESTS**

#### **Handler Tests**
bash
go test ./internal/handler/seller/ -v

#### **Service Tests**
bash
go test ./internal/service/seller/ -v

#### **Repository Tests**
bash
go test ./internal/repository/seller/ -v

#### **All Seller tests**
bash
go test ./internal/handler/seller/ ./internal/service/seller/ ./internal/repository/seller/ -v

---

## 📈 Coverage by Module

### **Locality Coverage**
bash
# Individual coverage by layer
go test ./internal/handler/locality/ -cover
go test ./internal/service/locality/ -cover
go test ./internal/repository/locality/ -cover

# Complete module coverage with Makefile
make cover  # Includes locality in the general report

### **Seller Coverage**
bash
# Individual coverage by layer
go test ./internal/handler/seller/ -cover
go test ./internal/service/seller/ -cover
go test ./internal/repository/seller/ -cover

# Complete module coverage with Makefile
make cover  # Includes seller in the general report

---

## 🎯 Specific Tests

### **Run specific test**
bash
# Example: run only a specific test
go test ./internal/service/seller/ -run TestSellerService_Create -v
go test ./internal/service/locality/ -run TestLocalityService_Create -v

### **Run tests matching pattern**
bash
# Run all Create tests
go test ./internal/service/seller/ -run ".*Create.*" -v

---

## 🔍 Quick Verification

### **Current Coverage Status**
- **Locality**: 100% on all layers ✅
- **Seller**: 100% on all layers ✅

### **Verification commands using Makefile**
bash
# Run all project tests
make test

# Generate complete coverage report
make cover-total

### **Direct command to verify only Locality & Seller**
bash
go test ./internal/handler/locality/ ./internal/service/locality/ ./internal/repository/locality/ ./internal/handler/seller/ ./internal/service/seller/ ./internal/repository/seller/ -v

---

## 📝 Important Notes

- **Makefile commands**: Use `make test`, `make cover`, `make cover-total`
- **HTML Coverage**: `make cover` opens the report in the browser
- **Total coverage**: `make cover-total` excludes mocks and test files
- Tests use **testify/require** for assertions
- Tests implemented with **table-driven testing** pattern


---

## 🚀 Quick Usage Examples

bash
# 1. Run all project tests (recommended)
make test

# 2. View coverage with nice HTML report
make cover

# 3. View total coverage without mocks or test files
make cover-total

# 4. Run only both modules tests (specific)
go test ./internal/handler/locality/ ./internal/service/locality/ ./internal/repository/locality/ ./internal/handler/seller/ ./internal/service/seller/ ./internal/repository/seller/ -v

# 5. Run only Service layer tests
go test ./internal/service/locality/ ./internal/service/seller/ -v

# 6. Run specific Seller Create test
go test ./internal/service/seller/ -run TestSellerService_Create -v

### **✅ Expected Results**
bash
# make test
PASS
ok      ProyectoFinal/internal/handler/locality    0.863s
ok      ProyectoFinal/internal/service/locality    0.845s  
ok      ProyectoFinal/internal/repository/locality 1.035s
ok      ProyectoFinal/internal/handler/seller      0.896s
ok      ProyectoFinal/internal/service/seller      0.999s
ok      ProyectoFinal/internal/repository/seller   1.173s

# make cover-total
total: (statements) 71.5%

### **🌟 Most Useful Commands**
1. **`make test`** - Quick verification of all tests
2. **`make cover`** - Visual coverage with HTML
3. **`make cover-total`** - Real project coverage (without mocks) 