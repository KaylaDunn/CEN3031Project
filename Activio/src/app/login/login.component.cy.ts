import { LoginComponent } from "./login.component"; // import the login component

// front end's first cypress test!
// sources:
    // https://docs.cypress.io/guides/component-testing/angular/quickstart
    // https://docs.cypress.io/guides/component-testing/angular/examples

// this is the "spec file" for Login Component
// "describe" and "it" allow us to group tests into sections by using method blocks

// "describe" block contains all our tests in a file
// takes 2 params: name of test suite & function that will execute the tests
describe('LoginComponent', () => {
    // each "it" block is a test
    // first param: brief description of spec
    // 2nd param: function that contains test code
    it('mounts', () => {
        // this method will mount our component into the test app so we can begin running tests against it
        cy.mount(LoginComponent)
    })

    // next would be to test the component's functionality

    // check that both buttons are enabled
    // source: https://www.webtips.dev/webtips/cypress/check-if-buttons-are-disabled-in-cypress#:~:text=To%20check%20if%20your%20buttons%20are%20disabled%20in%20Cypress%20or,enabled')%20assertions.
    it('buttons enabled', () => {
        cy.mount(LoginComponent)
        cy.get('button').should('be.enabled')
    })
})