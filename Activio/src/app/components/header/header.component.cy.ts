import { HeaderComponent } from './header.component';

describe('HeaderComponent', () => {
    it('mounts', () => {
        cy.mount(HeaderComponent)
    })

    it('buttons enabled', () => {
        cy.mount(HeaderComponent)
        cy.get('button').should('be.enabled')
    })
})