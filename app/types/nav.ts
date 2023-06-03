export interface NavItem {
  title: string
  href?: string
  disabled?: boolean
  external?: boolean
}

export interface CustomError {
  message: string
}

export interface SessionProps {
  roomName: string
  identity: string

}