import { Component } from '@angular/core';
import { MatListModule } from '@angular/material/list';
//import { ratings } from './exampleRating';
import { GetStatisticsResponse } from 'src/openapi-client/evaluationapi';
import { StatisticsService } from 'src/openapi-client/evaluationapi';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.css']
})
export class ListComponent {
  public ratings: GetStatisticsResponse | undefined;
  view: [number, number] = [300, 250];

  gradient: boolean = false;
  showLegend: boolean = false;
  showLabels: boolean = true;
  isDoughnut: boolean = false;

  progressColors = 
  [
    { name: "rated", value: '#0099ff' }, 
    { name: "unrated", value: '#c9c9c9' },
  ]

  resultColors =
  [
    { name: "positive", value: '#24a800' },
    { name: "negative", value: '#d61500' },
    { name: "neutral", value: '#fff700'}
  ]
  
  constructor(private readonly statisticService: StatisticsService) {
  }

  ngOnInit() {
    this.statisticService.getStatistics().subscribe(e => {
      this.ratings = e
     Object.assign(this, { e });
    })
    
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
