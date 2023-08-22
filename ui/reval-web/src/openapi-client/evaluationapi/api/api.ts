export * from './responseEvaluation.service';
import { ResponseEvaluationService } from './responseEvaluation.service';
export * from './statistics.service';
import { StatisticsService } from './statistics.service';
export const APIS = [ResponseEvaluationService, StatisticsService];
