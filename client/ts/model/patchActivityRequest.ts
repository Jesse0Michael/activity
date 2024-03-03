/**
 * Activity API
 * Activity Management API
 *
 * The version of the OpenAPI document: 0.0.3
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from './models';
import { Media } from './media';

export class PatchActivityRequest {
    /**
    * The source platform the activity is from
    */
    'source'?: string;
    /**
    * Unix timestamp (seconds) for when the item was published
    */
    'ts'?: number;
    /**
    * Permalink to the activity from it\'s source
    */
    'url'?: string;
    /**
    * Text content for the item (may contain HTML)
    */
    'body'?: string;
    /**
    * Array of media items (images, videos, etc...)
    */
    'media'?: Array<Media>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "source",
            "baseName": "source",
            "type": "string"
        },
        {
            "name": "ts",
            "baseName": "ts",
            "type": "number"
        },
        {
            "name": "url",
            "baseName": "url",
            "type": "string"
        },
        {
            "name": "body",
            "baseName": "body",
            "type": "string"
        },
        {
            "name": "media",
            "baseName": "media",
            "type": "Array<Media>"
        }    ];

    static getAttributeTypeMap() {
        return PatchActivityRequest.attributeTypeMap;
    }
}
