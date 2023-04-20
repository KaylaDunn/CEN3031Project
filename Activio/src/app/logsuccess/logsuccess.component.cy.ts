import { LogsuccessComponent } from './logsuccess.component';
import { HttpClientModule } from '@angular/common/http'
import { MatDialogModule, MatDialogRef, MAT_DIALOG_DATA } from '@angular/material/dialog'
import { AuthService } from '../auth.service';

describe('LogsuccessComponent', () => {
    beforeEach(() => {
        cy.mount(LogsuccessComponent, {
            imports: [HttpClientModule, MatDialogModule],
            providers: [
                { provide: MAT_DIALOG_DATA, useValue: {} },
                { provide: MatDialogRef, useValue: {} },
                { provide: AuthService, useValue: {} }
            ]
        })
    })

    it('buttons enabled', () => {
        cy.get('button').should('be.enabled')
    })
})