import { LogLevel } from '@nestjs/common';
import { registerAs } from '@nestjs/config';

export default registerAs('config', () => ({
    //
    // PORT: The port number this service is listening.
    //
    port: parseInt(process.env['PORT'], 10) || 3333,
    //
    // PORT: The port number this service is listening.
    //
    globalPrefix: process.env['GLOBAL_PREFIX'] || '/api',
    //
    // LOG_LEVEL: Like "error,warn" or "log,error,warn,debug,verbose".
    //
    logLevel: (process.env['LOG_LEVEL']?.split(',') || [
        'log',
        'error',
        'warn',
        'debug',
    ]) as LogLevel[],
    //
    // Redis
    //
    database: {
        host: process.env['DATABASE_HOST'] || 'localhost',
        port: parseInt(process.env['DATABASE_PORT'], 10) || 6379,
    },
}));
