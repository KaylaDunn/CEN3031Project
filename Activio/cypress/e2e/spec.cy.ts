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

  /*
  it('navigates to logged in page', () => {
    cy.visit('/login')
    cy.get('input[type="text"]').type('myUsername123')
    cy.get('input[type="password"]').type('myPassword123')
    cy.get('button[data-cy="b1"]').click()
    cy.url().should('eq', 'http://localhost:4200/logsuccess')
  })
  */

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
  
})
