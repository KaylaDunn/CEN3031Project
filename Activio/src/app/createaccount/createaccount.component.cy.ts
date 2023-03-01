import { CreateaccountComponent } from './createaccount.component';

describe('CreateaccountComponent', () => {
    it('mounts', () => {
        cy.mount(CreateaccountComponent)
    })

    it('buttons are clickable', () => {
        cy.mount(CreateaccountComponent)
        cy.get('button').should('be.enabled')
    })
})