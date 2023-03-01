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
})