import { NgModule } from '@angular/core';
import { AuthModule, AuthInterceptor } from 'angular-auth-oidc-client';
import { HTTP_INTERCEPTORS} from '@angular/common/http';

@NgModule({
    imports: [AuthModule.forRoot({
        config: {
              authority: 'https://th-b.eu.auth0.com',
              redirectUrl: window.location.origin,
              postLogoutRedirectUri: window.location.origin,
              clientId: 'FNitzlm8QQkjmjDmMG1m2pgHyOATo1xo',
              scope: 'openid profile email offline_access',
              responseType: 'code',
              silentRenew: true,
              useRefreshToken: true,
              renewTimeBeforeTokenExpiresInSeconds: 30,
              secureRoutes: ['http://localhost:8080/api'],
              customParamsAuthRequest: {
                audience: 'http://localhost:8080/api',
              },
          }
      })],
      providers: [
        { provide: HTTP_INTERCEPTORS, useClass: AuthInterceptor, multi: true },
    ],
    exports: [AuthModule],
})
export class AuthConfigModule {}
