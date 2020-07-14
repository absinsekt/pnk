import { linear } from 'svelte/easing';


export function spin(node, { idx, duration }) {
  return {
    duration,
    css: t => {
      const eased = linear(t);

      return (idx % 2) === 0
        ? `transform: rotate(${180 + eased * -180}deg); opacity: ${eased};`
        : `transform: rotate(${180 + eased * 180}deg); opacity: ${eased};`
    }
  };
}
