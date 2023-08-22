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


/**
 * Evaluation statistics.
 */
export interface Evaluations { 
    /**
     * number of neutral evaluations.
     */
    numNeutral: number;
    /**
     * Number of positive evaluations.
     */
    numPositive: number;
    /**
     * number of negative evaluations.
     */
    numNegative: number;
}

