import { Component } from '@angular/core';
import { ResponseEvaluationService } from '../../openapi-client/evaluationapi/api/responseEvaluation.service'
import { GetEvaluationResponse } from '../../openapi-client/evaluationapi/model/getEvaluationResponse'
import { PostEvaluationRequest } from 'src/openapi-client/evaluationapi';

@Component({
  selector: 'app-evaluation',
  templateUrl: './evaluation.component.html',
  styleUrls: ['./evaluation.component.sass']
})

export class EvaluationComponent {
  public evaluation : GetEvaluationResponse | undefined

  constructor(private readonly responseEvaluationService: ResponseEvaluationService) {}
  
  ngOnInit() {
    this.responseEvaluationService.getEvaluation().subscribe(e => {
      
      this.evaluation = e
      console.log("ngOnInit ")
      console.log(e)
      
    });
  }

  onNext(){

    var req : PostEvaluationRequest = {
      id: this.evaluation?.id || 0,
      evaluationResult: PostEvaluationRequest.EvaluationResultEnum.Neutral
    };

    this.sendAndGetNew(req as PostEvaluationRequest);
  }

  onPositive(){
    let req : PostEvaluationRequest = ({
      id: this.evaluation?.id || 0,
      evaluationResult: PostEvaluationRequest.EvaluationResultEnum.Positive
    });
    
    this.sendAndGetNew(req);
  }

  onNegative(){
    let req : PostEvaluationRequest = ({
      id: this.evaluation?.id || 0,
      evaluationResult: PostEvaluationRequest.EvaluationResultEnum.Negative
    });
   
    this.sendAndGetNew(req);
  }

  sendAndGetNew(req: PostEvaluationRequest){
    this.responseEvaluationService.postEvaluation(undefined, req).subscribe( _ => 
      this.responseEvaluationService.getEvaluation().subscribe(e => this.evaluation = e)
    );
  }
}
