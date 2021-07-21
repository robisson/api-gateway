import * as yaml from 'js-yaml'
import * as fs from 'fs';

const loadConfig = (configFile: string, encode: BufferEncoding = 'utf8') => {
    try {
        return yaml.load(fs.readFileSync(configFile, encode));
    } catch (e) {
        console.log(e);
    }
}

const configApi = loadConfig(__dirname + "/../api-gateway.yaml");

