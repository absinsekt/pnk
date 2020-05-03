import IMask from 'imask';


const initMasks = () => {
  const phones = Array.from(document.querySelectorAll('input[type="tel"]'));

  for (let p of phones) {
    IMask(p as HTMLElement, { mask: '+{7} (000) 000-00-00' });
  }
}

export { initMasks }
