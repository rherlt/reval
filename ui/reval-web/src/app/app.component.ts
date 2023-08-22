import { Component, OnInit} from '@angular/core';
import { OidcSecurityService } from 'angular-auth-oidc-client';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {
  title = 'reval-web';
  isAuthenticated = false
  constructor(public oidcSecurityService: OidcSecurityService) {}

  ngOnInit() {
    this.oidcSecurityService
      .checkAuth()
      .subscribe(({ isAuthenticated }) => {
        this.isAuthenticated = isAuthenticated;
      });
  }

  login(){
    console.log('login clicked');
    this.oidcSecurityService.authorize();
  }
}