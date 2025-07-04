# 🧪 Nimbus Demo Testing Guide

## 🎯 Demo Account Testing

### Demo Account Information
- **Demo Account Number**: `DEMO-1234567890`
- **Simplified ID for Testing**: `1234567890`
- **Purpose**: Send money to this account to see automatic responses
- **Auto-Response**: You'll receive 2 demo transactions back (50% and 25% of sent amount)

### How to Test the Demo
1. **Create Account & Verify Email**
   - Register a new account
   - Check your email for verification link
   - Click to verify (receives $100 welcome credits)

2. **Use Welcome Credits**
   - Initial account gets $100 USD
   - Create additional accounts in different currencies
   - EUR accounts get €95 welcome credits
   - INR accounts get ₹8300 welcome credits

3. **Test Demo Transactions**
   - Send any amount to account: `1234567890`
   - Wait 10 seconds for first auto-response (50% return)
   - Wait 5 more seconds for second auto-response (25% return)
   - Check your balance and transaction history

### Test Scenarios

| Test Case | Action | Expected Result |
|-----------|--------|----------------|
| **Email Verification** | Register new account | Receive verification email, $100 credit after verification |
| **Currency Conversion** | Create EUR account | Welcome amount shows as €95 (based on exchange rate) |
| **Currency Conversion** | Create INR account | Welcome amount shows as ₹8300 (based on exchange rate) |
| **Demo Transaction** | Send $10 to demo account | Receive $5 back (50%), then $2.50 (25%) |
| **Real-time Updates** | Make any transfer | Balance updates immediately |
| **Account Numbers** | Use simplified demo ID | System maps `1234567890` → `DEMO-1234567890` |

## 🔧 Technical Features Implemented

### ✅ Fixed Issues
1. **Transfer Account Lookup**
   - Added `GetAccountByNumber` SQL function
   - Enhanced transfer logic to handle both account IDs and account numbers
   - Special mapping for demo account: `1234567890` → `DEMO-1234567890`

2. **Currency-Specific Welcome Credits**
   - USD: $100.00 (10,000 cents)
   - EUR: €95.00 (9,500 cents)
   - INR: ₹8300.00 (830,000 paisa)

3. **Demo Response System**
   - Automatic detection of transfers to demo account
   - Async processing with Redis/Asynq
   - Two-stage response: 50% after 10s, 25% after 15s

4. **Frontend Improvements**
   - Transfer form accepts both numeric IDs and account numbers
   - Better error messages and user guidance
   - Updated help text with demo account number

### 🏗️ Architecture
- **Backend**: Go + PostgreSQL + Redis
- **Frontend**: React + Bootstrap
- **Queue System**: Asynq for demo responses
- **Security**: JWT authentication + email verification

## 🎮 For Recruiters & Testers

### Quick Demo (5 minutes)
1. Go to the website
2. Register with your email
3. Verify email (check spam folder)
4. Create a few accounts in different currencies
5. Send $20 to demo account: `1234567890`
6. Watch the magic happen! 🎩✨

### What This Demonstrates
- **Full-Stack Development**: Go backend + React frontend
- **Database Design**: PostgreSQL with proper relationships and transactions
- **Security**: JWT auth, email verification, input validation
- **Real-time Processing**: Instant balance updates, async demo responses
- **Multi-Currency Support**: Proper currency conversion and handling
- **Error Handling**: Comprehensive error management
- **UI/UX**: Modern, responsive design
- **Testing**: Automated demo system for easy evaluation

## 📞 Contact
- **Email**: singhom2003.os@gmail.com
- **Phone**: +91 8810967896

---
*Built with ❤️ for secure banking by Om Singh*
