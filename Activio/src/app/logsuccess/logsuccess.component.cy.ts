import { LogsuccessComponent } from './logsuccess.component';

describe('LogsuccessComponent', () => {
    it('mounts', () => {
        cy.mount(LogsuccessComponent)
    })

    it('buttons enabled', () => {
        cy.mount(LogsuccessComponent)
        cy.get('button').should('be.enabled')
    })
})