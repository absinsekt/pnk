import { isMobile } from './core/detect';
import { initApps } from './applications';

if (!isMobile()) {
  document.body.classList.add('desktop');
}

initApps();
