import { Component } from '@angular/core';
import { MatListModule } from '@angular/material/list';
import { evaluations } from './data';
import { ratings } from './exampleRating';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.css']
})
export class ListComponent {
  ratings: any[] | undefined;
  view: [number, number] = [400, 400];

  gradient: boolean = false;
  showLegend: boolean = false;
  showLabels: boolean = true;
  isDoughnut: boolean = true;

  customColors = 
  [
    { name: "negative", value: '#eb4034' }, 
    { name: "unanswered", value: '#fae714' },
    { name: "positive", value: '#19bf19'},
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
