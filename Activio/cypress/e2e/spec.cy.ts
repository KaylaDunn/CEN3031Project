// reference: https://testing-angular.com/end-to-end-testing/#interacting-with-elements

describe('My First Test', () => {

  it('Visits the initial project page', () => {
    cy.visit('/')
    cy.contains('Welcome to Activio!')
  })

  it('navigates to Login page', () => {
    cy.visit('/')
    cy.get('button[type="loginBtn"]').click()
    cy.url().should('eq', 'http://localhost:4200/login')
  })

  it('navigates to logged in page', () => {
    cy.visit('/login')
    cy.get('input[type="text"]').type('myUsername123')
    cy.get('input[type="password"]').type('myPassword123')
    cy.get('button[type="button"]').click()
    cy.url().should('eq', 'http://localhost:4200/logsuccess')
  })

  it('navigates to Signup page', () => {
    cy.visit('/login')
    cy.get('button[type="signup"]').click()
    cy.url().should('eq', 'http://localhost:4200/createaccount')
  })

})
