import { isSet } from './objects';

export function getWgtGroup(target: HTMLElement): string {
  let wgt = target.parentElement;

  for (let i = 0; i < 10; i++) {
    if (isSet(wgt) && wgt.classList.contains('pnk-wgt')) {
      return wgt.getAttribute('data-group') || null;
    }

    if (isSet(wgt) && wgt.parentElement) {
      wgt = wgt.parentElement;
    } else {
      return null;
    }
  }

  return null;
}
