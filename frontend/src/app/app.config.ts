import {APP_INITIALIZER, ApplicationConfig} from '@angular/core';
import {provideRouter} from '@angular/router';

import {routes} from './app.routes';
import {provideAnimations} from "@angular/platform-browser/animations";
import {PrimeNGConfig} from "primeng/api";
import {provideHttpClient} from "@angular/common/http";

const initializeAppPrimeNGConfigFactory = (primeConfig: PrimeNGConfig) => () => {
  primeConfig.ripple = true;
};

export const appConfig: ApplicationConfig = {
  providers: [
    provideHttpClient(),
    provideRouter(routes),
    provideAnimations(),
    {
      provide: APP_INITIALIZER,
      useFactory: initializeAppPrimeNGConfigFactory,
      multi: true,
      deps: [PrimeNGConfig],
    },]

};

