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

  it('navigates to Signup page', () => {
    cy.visit('/login')
    cy.get('button[type="signup"]').click()
    cy.url().should('eq', 'http://localhost:4200/createaccount')
  })

})
