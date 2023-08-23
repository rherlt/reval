import { Component } from '@angular/core';
import { MatListModule } from '@angular/material/list';
import { evaluations } from './data';
import { ratings } from './exampleRating';
import { GetStatisticsResponse } from 'src/openapi-client/evaluationapi';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.css']
})
export class ListComponent {
  public ratings: GetStatisticsResponse | undefined;
  view: [number, number] = [400, 400];

  gradient: boolean = false;
  showLegend: boolean = false;
  showLabels: boolean = true;
  isDoughnut: boolean = true;

  customColors = 
  [
    { name: "rated", value: '#0099ff' }, 
    { name: "unrated", value: '#c9c9c9' },
  ]
  
  constructor() {
    Object.assign(this, { ratings });
  }

  onSelect(data: any): void {
    console.log('Item clicked', JSON.parse(JSON.stringify(data)));
  }

  onActivate(data: any): void {
    console.log('Activate', JSON.parse(JSON.stringify(data)));
  }

  onDeactivate(data: any): void {
    console.log('Deactivate', JSON.parse(JSON.stringify(data)));
  }
}
