export const inColor = '#4caf50';
export const outColor = '#ff1744';

export function toggleColor(currentColor: string): string {
    return currentColor === inColor ? outColor : inColor;
}