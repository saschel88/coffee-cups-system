# Coffee Cups System - Development Rules

## 🎯 Project Overview
This document defines the rules, standards, and best practices for the Coffee Cups System project - a Telegram-based coffee consumption tracking system with fair cost distribution.

## 📋 Table of Contents
1. [General Rules](#general-rules)
2. [Code Standards](#code-standards)
3. [API Development](#api-development)
4. [Database Rules](#database-rules)
5. [Testing Rules](#testing-rules)
6. [Documentation Rules](#documentation-rules)
7. [Git Workflow](#git-workflow)
8. [Security Rules](#security-rules)
9. [Performance Rules](#performance-rules)
10. [Deployment Rules](#deployment-rules)
11. [Telegram Bot Rules](#telegram-bot-rules)
12. [Business Logic Rules](#business-logic-rules)

---

## 🏗️ General Rules

### 1.1 Project Structure
- **MUST** follow Clean Architecture principles
- **MUST** use Domain-Driven Design (DDD) patterns
- **MUST** separate concerns into distinct layers:
  - `cmd/` - Application entry points (server, migrate, server-no-bot)
  - `internal/` - Private application code
    - `config/` - Configuration management
    - `database/` - Database connection and setup
    - `handlers/` - HTTP request handlers
    - `logger/` - Logging utilities
    - `models/` - Data models (User, Box, CoffeeLog, Payment)
    - `services/` - Business logic services
    - `server/` - HTTP server setup
    - `telegram/` - Telegram bot implementation
  - `api/` - API specifications (swagger.yaml)
  - `docs/` - Documentation
  - `scripts/` - Build and deployment scripts
  - `configs/` - Configuration files

### 1.2 Naming Conventions
- **MUST** use descriptive, self-documenting names
- **MUST** use camelCase for Go variables and functions
- **MUST** use PascalCase for Go types and public functions
- **MUST** use snake_case for database fields and API endpoints
- **MUST** use kebab-case for file names and directories

### 1.3 Error Handling
- **MUST** handle all errors explicitly
- **MUST** use structured logging for errors
- **MUST** return meaningful error messages to users
- **MUST NOT** ignore errors with `_`

---

## 💻 Code Standards

### 2.1 Go Code Rules
```go
// ✅ GOOD
func (s *UserService) GetUserByID(id uint32) (*models.User, error) {
    if id == 0 {
        return nil, errors.New("invalid user ID")
    }
    // implementation
}

// ❌ BAD
func GetUser(id uint32) *models.User {
    // no error handling, no validation
}
```

### 2.2 Code Organization
- **MUST** have one main function per file in `cmd/`
- **MUST** group related functionality in packages
- **MUST** use interfaces for external dependencies
- **MUST** keep functions under 50 lines
- **MUST** keep files under 300 lines

### 2.3 Comments and Documentation
- **MUST** document all public functions
- **MUST** use `//` for single-line comments
- **MUST** use `/* */` for multi-line comments
- **MUST** explain complex business logic

```go
// CreateUser creates a new user in the system
// Returns the created user or an error if creation fails
func (s *UserService) CreateUser(req *CreateUserRequest) (*models.User, error) {
    // implementation
}
```

---

## 🌐 API Development

### 3.1 Swagger-First Approach
- **MUST** define API in `api/swagger.yaml` first
- **MUST** generate code from OpenAPI specification
- **MUST** validate all requests against the schema
- **MUST** return consistent response formats

### 3.2 API Endpoints
- **MUST** use RESTful conventions
- **MUST** use HTTP status codes correctly:
  - `200` - Success
  - `201` - Created
  - `400` - Bad Request
  - `404` - Not Found
  - `500` - Internal Server Error
- **MUST** use proper HTTP methods:
  - `GET` - Retrieve data
  - `POST` - Create new resources
  - `PUT` - Update existing resources
  - `DELETE` - Remove resources

### 3.3 Request/Response Format
```json
// ✅ GOOD - Consistent response format
{
  "data": { ... },
  "message": "Success",
  "timestamp": "2025-01-01T00:00:00Z"
}

// ❌ BAD - Inconsistent format
{
  "user": { ... },
  "status": "ok"
}
```

### 3.4 API Versioning
- **MUST** version all APIs (`/api/v1/`)
- **MUST** maintain backward compatibility
- **MUST** deprecate old versions gracefully

---

## 🗄️ Database Rules

### 4.1 Schema Design
- **MUST** use meaningful table and column names
- **MUST** include `created_at` and `updated_at` timestamps
- **MUST** use appropriate data types
- **MUST** add indexes for frequently queried columns
- **MUST** use foreign key constraints

### 4.2 Migrations
- **MUST** create migration files for schema changes
- **MUST** make migrations reversible
- **MUST** test migrations on development data
- **MUST** backup production data before migrations

### 4.3 Queries
- **MUST** use parameterized queries to prevent SQL injection
- **MUST** use transactions for multi-table operations
- **MUST** avoid N+1 query problems
- **MUST** use appropriate indexes

```go
// ✅ GOOD
db.Where("user_id = ?", userID).Find(&logs)

// ❌ BAD
db.Raw(fmt.Sprintf("SELECT * FROM logs WHERE user_id = %d", userID))
```

---

## 🧪 Testing Rules

### 5.1 Test Coverage
- **MUST** achieve minimum 80% code coverage
- **MUST** write unit tests for all business logic
- **MUST** write integration tests for API endpoints
- **MUST** write tests for error scenarios

### 5.2 Test Structure
```go
func TestUserService_CreateUser(t *testing.T) {
    // Arrange
    service := NewUserService(mockDB, mockLogger)
    req := &CreateUserRequest{
        TelegramID: 123456789,
        FirstName:  "John",
    }
    
    // Act
    user, err := service.CreateUser(req)
    
    // Assert
    assert.NoError(t, err)
    assert.Equal(t, "John", user.FirstName)
}
```

### 5.3 Test Data
- **MUST** use test fixtures for consistent data
- **MUST** clean up test data after each test
- **MUST** use factories for creating test objects
- **MUST NOT** use production data in tests

---

## 📚 Documentation Rules

### 6.1 API Documentation
- **MUST** document all API endpoints in Swagger
- **MUST** provide example requests and responses
- **MUST** document error codes and messages
- **MUST** keep documentation up-to-date

### 6.2 Code Documentation
- **MUST** document all public APIs
- **MUST** provide usage examples
- **MUST** document configuration options
- **MUST** document deployment procedures

### 6.3 README Files
- **MUST** have a main README.md in project root
- **MUST** include setup and installation instructions
- **MUST** include usage examples
- **MUST** include contribution guidelines

---

## 🔄 Git Workflow

### 7.1 Branch Strategy
- **MUST** use feature branches for new development
- **MUST** use `main` branch for production-ready code
- **MUST** use descriptive branch names: `feature/user-authentication`
- **MUST** delete feature branches after merging

### 7.2 Commit Messages
- **MUST** use conventional commit format:
  - `feat:` - New features
  - `fix:` - Bug fixes
  - `docs:` - Documentation changes
  - `style:` - Code style changes
  - `refactor:` - Code refactoring
  - `test:` - Test additions/changes
  - `chore:` - Build process or auxiliary tool changes

### 7.3 Pull Requests
- **MUST** create pull requests for all changes
- **MUST** include description of changes
- **MUST** link related issues
- **MUST** request code review
- **MUST** ensure all tests pass

---

## 🔒 Security Rules

### 8.1 Input Validation
- **MUST** validate all user inputs
- **MUST** sanitize data before database operations
- **MUST** use parameterized queries
- **MUST** validate file uploads

### 8.2 Authentication & Authorization
- **MUST** implement proper authentication
- **MUST** use secure session management
- **MUST** implement role-based access control
- **MUST** log security events

### 8.3 Data Protection
- **MUST** encrypt sensitive data
- **MUST** use HTTPS in production
- **MUST** implement rate limiting
- **MUST** validate JWT tokens

---

## ⚡ Performance Rules

### 9.1 Database Performance
- **MUST** use appropriate indexes
- **MUST** avoid N+1 queries
- **MUST** use connection pooling
- **MUST** monitor query performance

### 9.2 API Performance
- **MUST** implement response caching where appropriate
- **MUST** use pagination for large datasets
- **MUST** implement rate limiting
- **MUST** monitor API response times

### 9.3 Resource Management
- **MUST** close database connections properly
- **MUST** handle memory efficiently
- **MUST** implement graceful shutdown
- **MUST** monitor resource usage

---

## 🚀 Deployment Rules

### 10.1 Environment Configuration
- **MUST** use environment variables for configuration
- **MUST** have separate configs for dev/staging/prod
- **MUST** never commit secrets to version control
- **MUST** use configuration management tools

### 10.2 Docker Rules
- **MUST** use multi-stage builds
- **MUST** use specific image tags
- **MUST** run as non-root user
- **MUST** use health checks

### 10.3 Monitoring
- **MUST** implement health checks
- **MUST** log application metrics
- **MUST** monitor error rates
- **MUST** set up alerts for critical issues

---

## 📊 Quality Gates

### 11.1 Code Quality
- **MUST** pass all linter checks
- **MUST** achieve minimum 80% test coverage
- **MUST** pass security scans
- **MUST** follow coding standards

### 11.2 Performance Requirements
- **MUST** respond to API requests within 200ms
- **MUST** handle at least 100 concurrent users
- **MUST** maintain 99.9% uptime
- **MUST** use less than 512MB memory per instance

---

## 🚫 Anti-Patterns to Avoid

### 12.1 Code Anti-Patterns
- **MUST NOT** use global variables
- **MUST NOT** ignore errors
- **MUST NOT** use magic numbers
- **MUST NOT** create god objects/functions

### 12.2 API Anti-Patterns
- **MUST NOT** expose internal implementation details
- **MUST NOT** use inconsistent response formats
- **MUST NOT** return sensitive data in responses
- **MUST NOT** use GET for state-changing operations

### 12.3 Database Anti-Patterns
- **MUST NOT** use SELECT * in production
- **MUST NOT** create tables without proper indexes
- **MUST NOT** use string types for numeric data
- **MUST NOT** store passwords in plain text

---

## 📝 Enforcement

### 13.1 Automated Checks
- **MUST** run linters in CI/CD pipeline
- **MUST** run tests automatically
- **MUST** check code coverage
- **MUST** validate API schemas

### 13.2 Code Reviews
- **MUST** review all code changes
- **MUST** check for rule compliance
- **MUST** verify test coverage
- **MUST** ensure documentation updates

---

## 📞 Contact & Support

For questions about these rules or to suggest improvements:
- Create an issue in the project repository
- Contact the development team
- Review the project documentation

---

## 🤖 Telegram Bot Rules

### 11.1 Bot Implementation
- **MUST** handle all Telegram API errors gracefully
- **MUST** implement proper command validation
- **MUST** use structured logging for bot interactions
- **MUST** handle user registration automatically
- **MUST** provide clear error messages to users

### 11.2 Command Structure
- **MUST** use consistent command format: `/command [parameters]`
- **MUST** implement help command with all available commands
- **MUST** validate all user inputs from Telegram
- **MUST** handle malformed commands gracefully

### 11.3 Bot Security
- **MUST** validate Telegram user IDs
- **MUST** implement rate limiting for bot commands
- **MUST** log all bot interactions for audit
- **MUST** handle bot token security properly

---

## 💼 Business Logic Rules

### 12.1 Coffee Tracking Logic
- **MUST** validate box existence before logging coffee
- **MUST** check box capacity before allowing consumption
- **MUST** calculate fair share based on actual consumption
- **MUST** handle edge cases (empty boxes, invalid users)

### 12.2 Payment Calculation
- **MUST** calculate payments based on actual consumption
- **MUST** handle partial payments correctly
- **MUST** track payment status accurately
- **MUST** provide payment history and summaries

### 12.3 Data Integrity
- **MUST** maintain referential integrity between entities
- **MUST** handle soft deletes properly
- **MUST** validate business rules at service layer
- **MUST** prevent orphaned records

### 12.4 Service Layer Rules
- **MUST** implement all business logic in services
- **MUST** use dependency injection for services
- **MUST** handle transactions properly
- **MUST** provide clear error messages for business rule violations

---

## 📊 Business Layer Description (Product Manager View)

### Core Business Value Proposition
The Coffee Cups System solves the common office problem of fair cost distribution for shared coffee consumption. It eliminates manual tracking, reduces disputes, and ensures everyone pays their fair share based on actual usage.

### Key Business Entities

#### 1. **Users** 👥
- **Purpose**: Track who is consuming coffee
- **Key Data**: Telegram ID, name, activity status
- **Business Rules**: 
  - Users are automatically registered via Telegram
  - Users can be deactivated but not deleted
  - Each user has a unique Telegram ID

#### 2. **Coffee Boxes** 📦
- **Purpose**: Represent a physical coffee capsule package
- **Key Data**: Name, total cups, price, creator, active status
- **Business Rules**:
  - Only active boxes can be consumed from
  - Box creator is responsible for initial cost
  - Box capacity is fixed and cannot be exceeded

#### 3. **Coffee Logs** ☕
- **Purpose**: Track individual coffee consumption
- **Key Data**: User, box, timestamp
- **Business Rules**:
  - Each log represents one cup consumed
  - Cannot log coffee from empty boxes
  - Timestamp is automatically set

#### 4. **Payments** 💰
- **Purpose**: Track who owes what for coffee consumption
- **Key Data**: User, box, amount, payment status
- **Business Rules**:
  - Payments are calculated based on actual consumption
  - Each user pays their proportional share
  - Payments can be marked as paid

### Business Workflows

#### 1. **Coffee Consumption Workflow**
1. User sends `/coffee <box_id>` to Telegram bot
2. System validates box exists and has remaining cups
3. System logs the consumption
4. System updates payment calculations
5. User receives confirmation with remaining cups

#### 2. **Box Management Workflow**
1. User creates a new box with name, cups, and price
2. System calculates cost per cup
3. System tracks consumption as users log coffee
4. System calculates each user's share based on consumption
5. Users can view their payment obligations

#### 3. **Payment Settlement Workflow**
1. System calculates each user's total owed amount
2. Users can view their payment status
3. Users can mark payments as paid
4. System tracks payment history

### Business Metrics & KPIs

#### 1. **Usage Metrics**
- Total coffee consumed per user
- Average consumption per user per period
- Most popular coffee boxes
- Consumption patterns by time/date

#### 2. **Financial Metrics**
- Total revenue from coffee sales
- Average cost per cup
- Payment collection rates
- Outstanding payment amounts

#### 3. **System Health Metrics**
- Active users count
- Active boxes count
- System uptime
- API response times

### Business Rules & Constraints

#### 1. **Consumption Rules**
- Users cannot consume more cups than available in a box
- Each consumption must be logged with a valid user and box
- Consumption timestamps are immutable

#### 2. **Payment Rules**
- Payment amounts are calculated proportionally based on consumption
- Users only pay for coffee they actually consumed
- Payment status can be updated but not deleted

#### 3. **Data Integrity Rules**
- Users cannot be deleted if they have consumption logs
- Boxes cannot be deleted if they have consumption logs
- All financial calculations must be auditable

### Value Proposition for Different Stakeholders

#### 1. **For Office Managers**
- Eliminates manual coffee cost tracking
- Reduces disputes over shared costs
- Provides clear financial reporting
- Automates fair cost distribution

#### 2. **For Employees**
- Simple Telegram-based interface
- Transparent cost calculation
- Easy payment tracking
- No manual record keeping

#### 3. **For Finance Teams**
- Automated cost allocation
- Detailed consumption reports
- Payment tracking and reconciliation
- Audit trail for all transactions

### Future Business Opportunities

#### 1. **Expansion Features**
- Multi-office support
- Advanced analytics and reporting
- Integration with expense management systems
- Mobile app for non-Telegram users

#### 2. **Monetization Opportunities**
- Premium features for larger organizations
- Analytics and reporting services
- Integration with corporate systems
- White-label solutions

---

**Last Updated:** 2025-01-04  
**Version:** 1.1.0  
**Maintainer:** Coffee Cups System Team
