export const WeekdaysShort = ['пн', 'вт', 'ср', 'чт', 'пт', 'сб', 'вс'];
export const WeekdaysFull = ['понедельник', 'вторник', 'среда', 'четверг', 'пятница', 'суббота', 'воскресенье'];
export const MonthsFull = ['январь', 'февраль', 'март', 'апрель', 'май', 'июнь', 'июль', 'август', 'сентябрь', 'октябрь', 'ноябрь', 'декабрь'];

export const enum DateEquality {
  Day = 0,
  Month = 1,
  Year = 2,
}

export function getStartOfMonth(now: Date): Date {
  const result = new Date(now);

  result.setDate(1);

  return result;
}

export function isEqualDate(date1: Date, date2: Date, eqType:DateEquality = DateEquality.Day): boolean {
  if (date1.getFullYear() !== date2.getFullYear()) return false;
  if (eqType === DateEquality.Year) return true;

  if (date1.getMonth() !== date2.getMonth()) return false;
  if (eqType === DateEquality.Month) return true;

  return date1.getDate() === date2.getDate();
}

export function getOffsetDate(date: Date, days: number, months: number, years: number) {
  const result = new Date(date);

  result.setDate(result.getDate() + days);
  result.setMonth(result.getMonth() + months);
  result.setFullYear(result.getFullYear() + years);

  return result;
}
