import { TestBed } from '@angular/core/testing';
import { PostComponent } from './post.component'
import { HttpClientTestingModule } from '@angular/common/http/testing';
import { AuthService } from '../auth.service';

describe('PostComponent', () => {

  beforeEach(() => {
    TestBed.configureTestingModule({
      imports: [HttpClientTestingModule],
      providers: [AuthService]
    }).compileComponents();

    cy.mount(PostComponent);
  });
  

  it('buttons enabled except post', () => {
    cy.get('button[type="homeBtn"]').should('be.enabled')
    cy.get('button[type="btnUp"]').should('be.enabled')
  })

  it('fills inputs', () => {
    cy.get('input[name="loc"]').type('Gainesville')
    cy.get('input[name="act"]').type('Getting Krishna')
  })

  // TODO:
  // it('post button disabled until input', () => {
  //   cy.mount(PostComponent)
  //   cy.get('button[name="postBtn"]').should('not.be.enabled')
  // })
})