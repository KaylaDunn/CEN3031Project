import { WelcomeComponent } from './welcome.component';
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { HttpClientModule } from '@angular/common/http';
import { TestBed } from '@angular/core/testing';
import { AuthService } from '../auth.service';

describe('WelcomeComponent', () => {
    beforeEach(() => {
        TestBed.configureTestingModule({
            imports: [HttpClientTestingModule, HttpClientModule],
            providers: [
                { provide: AuthService }
            ]
        }).compileComponents();

        cy.mount(WelcomeComponent);
    });

    it('buttons enabled', () => {
        cy.get('button').should('be.enabled')
    })
})