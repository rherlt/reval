/**
 * Evaluation API
 * This API is for evaluating responses from a response generator.
 *
 * The version of the OpenAPI document: 0.4
 * Contact: mail@rherlt.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


export interface RatingScore { 
    /**
     * Minimum of rating Score.
     */
    min: number;
    /**
     * Rating score of the scenario.
     */
    value: number;
    /**
     * Maximum of rating Score.
     */
    max: number;
}

