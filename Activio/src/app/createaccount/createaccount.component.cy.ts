import { TestBed } from '@angular/core/testing';
import { CreateaccountComponent } from './createaccount.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { AuthService } from '../auth.service';

describe('CreateaccountComponent', () => {

    beforeEach(() => TestBed.configureTestingModule({
        imports: [HttpClientTestingModule], 
        providers: [AuthService]
      }));

    it('mounts', () => {
        cy.mount(CreateaccountComponent)
    })

    it('buttons are clickable', () => {
        cy.mount(CreateaccountComponent)
        cy.get('button').should('be.enabled')
    })

    it('types input fields', () => {
        cy.mount(CreateaccountComponent)
        cy.get('input[name="loc"]').type('Gainesville')
        // cy.get('input[type="user"]').type('myUsername')
        // cy.get('input[type="password1"]').type('myPassword')
        // cy.get('input[type="password2"]').type('myPassword')
    })

    it('creates account', () => {
        // TODO: check if button will save data if all the fields are occupied

    })
})