import { isMobile } from '../pnk/core/detect';
import { initApps } from './applications';

if (!isMobile()) {
  document.body.classList.add('desktop');
}

initApps();
