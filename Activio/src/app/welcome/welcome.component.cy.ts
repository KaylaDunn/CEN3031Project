import { WelcomeComponent } from './welcome.component';
describe('WelcomeComponent', () => {
    it('mounts', () => {
        cy.mount(WelcomeComponent)
    })

    it('buttons enabled', () => {
        cy.mount(WelcomeComponent)
        cy.get('button').should('be.enabled')
    })
})