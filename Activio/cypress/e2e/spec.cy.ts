import { CreateaccountComponent } from './../../src/app/createaccount/createaccount.component';
// reference: https://testing-angular.com/end-to-end-testing/#interacting-with-elements

describe('Activio End-to-End', () => {

  // welcome
  it('Visits the initial project page', () => {
    cy.visit('/')
    cy.contains('Welcome to Activio!')
  })

  it('navigates to Login page', () => {
    cy.visit('/')
    cy.get('button[type="logBtn"]').click()
    cy.url().should('eq', 'http://localhost:4200/login')
  })

  it('navigates to logged in page', () => {
    cy.visit('/login')
    cy.get('input[type="text"]').type('b@b.com')
    cy.get('input[type="password"]').type('b123')
    cy.get('button[data-cy="b1"]').click()
    cy.url().should('eq', 'http://localhost:4200/logsuccess')
  })

  it('does not log in with invalid credentials', () => {
    cy.visit('/login')
    cy.get('input[type="text"]').type('fake@email.com')
    cy.get('input[type="password"]').type('fakePassword')
    cy.get('button[data-cy="b1"]').click()
    cy.url().should('eq', 'http://localhost:4200/login')
  })

  it('navigates to Signup page', () => {
    cy.visit('/login')
    cy.get('button[data-cy="b2"]').click()
    cy.url().should('eq', 'http://localhost:4200/createaccount')
  })

  it('logged in user can access post', () => {
    cy.visit('/logsuccess')
    cy.get('button[type="postBtn"]').click()
    cy.url().should('eq', 'http://localhost:4200/post')
  })

  it('post navigates to home', () => {
    cy.visit('/post')
    cy.get('button[type="homeBtn"]').click()
    cy.url().should('eq', 'http://localhost:4200/logsuccess')
  })

  it('createaccount navigates to verification', () => {
    cy.visit('/createaccount')
    cy.get('input[type="firstName"]').type('First')
    cy.get('input[type="lastName"]').type('Last')
    cy.get('input[type="email"]').type('this@email.com')
    cy.get('input[type="user"]').type('myUser')
    cy.get('input[type="birth"]').type('01/01/2001')
    cy.get('input[type="phone"]').type('(000)-000-000')
    cy.get('input[type="password1"]').type('pass1')
    cy.get('input[type="password2"]').type('pass1')
    cy.get('button[data-cy="signup"]').click()
    cy.url().should('eq', 'http://localhost:4200/create-account-verif')
  })

  it('verification navigates back to login', () => {
    cy.visit('/create-account-verif')
    cy.wait(6000)
    cy.url().should('eq', 'http://localhost:4200/login')
  })

  it('profile navigates to userProfile', () => {
    cy.visit('/logsuccess')
    cy.get('button[type="ProfileBttn"]').click()
    cy.url().should('eq', 'http://localhost:4200/userProfile')
  })

  it('edit account goes to edit account from userProfile', () => {
    cy.visit('/userProfile')
    cy.get('button[data-cy="editBtn"]').click()
    cy.url().should('eq', 'http://localhost:4200/editAccount')
  })

  it('user can log out', () => {
    cy.visit('/logsuccess')
    cy.get('button[type="LogBtn"]').click()
    cy.url().should('eq', 'http://localhost:4200/logout')
  })

  it('logout redirects to main', () => {
    cy.visit('/logout')
    cy.wait(6000)
    cy.url().should('eq', 'http://localhost:4200/')
  })
  
  it('post home button goes home', () => {
    cy.visit('/post')
    cy.get('button[type="homeBtn"]').click()
    cy.url().should('eq', 'http://localhost:4200/logsuccess')
  })
})
