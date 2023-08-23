/**
 * Evaluation API
 * This API is for evaluating responses from a response generator.
 *
 * The version of the OpenAPI document: 0.2
 * Contact: mail@rherlt.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */
import { NameValuePair } from './nameValuePair';
import { StatisticsScenarioRatingScore } from './statisticsScenarioRatingScore';


/**
 * Statistics per Scenario.
 */
export interface StatisticsScenario { 
    /**
     * Unique id of the scenario.
     */
    id?: string;
    /**
     * The name of the Scenario
     */
    name?: string;
    /**
     * The description of the Scenario.
     */
    desc?: string;
    /**
     * The amount of questions and response evaluated in this scenario.
     */
    totalResponseCount?: number;
    progressStatistics?: Array<NameValuePair>;
    resultStatistics?: Array<NameValuePair>;
    ratingScore?: StatisticsScenarioRatingScore;
}
