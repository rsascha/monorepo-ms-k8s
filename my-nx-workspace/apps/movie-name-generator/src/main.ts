import { Logger } from '@nestjs/common';
import { NestFactory } from '@nestjs/core';

import { AppModule } from './app/app.module';
import { environment } from './environments/environment';

async function bootstrap() {
    const config = environment.config;
    const logger = new Logger('main -> bootstrap()');
    logger.verbose(`Starting with config: ${JSON.stringify(config)}`);

    const app = await NestFactory.create(AppModule, {
        logger: config.logLevel,
    });

    app.setGlobalPrefix(config.globalPrefix);

    app.listen(config.port).then(() => {
        app.getUrl().then((url) => {
            logger.log(`Listening at ${url}`);
        });
    });
}

bootstrap();
