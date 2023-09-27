import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { HttpClientModule } from '@angular/common/http';
import { ApiModule } from '../openapi-client/evaluationapi/api.module';
import { MatIconModule } from '@angular/material/icon';
import { MatDividerModule } from '@angular/material/divider';
import { MatButtonModule } from '@angular/material/button';
import { MatCardModule } from '@angular/material/card';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatBadgeModule } from '@angular/material/badge';
import { EvaluationComponent } from './evaluation/evaluation.component';
import { AuthConfigModule } from './auth/auth-config.module';
import { UserprofileComponent } from './userprofile/userprofile.component';
import { StatisticsComponent } from './statistics/statistics.component';
import { MatListModule } from '@angular/material/list';
import { NgxChartsModule } from '@swimlane/ngx-charts';
import { Configuration } from 'src/openapi-client/evaluationapi';

@NgModule({
  declarations: [
    AppComponent,
    EvaluationComponent,
    UserprofileComponent,
    StatisticsComponent
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    BrowserAnimationsModule,
    HttpClientModule,
    ApiModule.forRoot(() => {
      return new Configuration({
        basePath: 'https://reval.th-b.com/api',
      })}),
    MatIconModule,
    MatDividerModule,
    MatButtonModule,
    MatCardModule,
    MatGridListModule,
    MatBadgeModule,
    AuthConfigModule,
    MatListModule,
    NgxChartsModule,
  ],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }
