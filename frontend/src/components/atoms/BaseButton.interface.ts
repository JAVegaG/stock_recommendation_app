export interface Props {
  variant?: 'primary' | 'secondary' | 'success' | 'danger' | 'warning' | 'ghost' | 'outline'
  size?: 'xs' | 'sm' | 'md' | 'lg' | 'xl'
  disabled?: boolean
  loading?: boolean
  loadingText?: string
  iconLeft?: any
  iconRight?: any
  fullWidth?: boolean
  rounded?: boolean
}
