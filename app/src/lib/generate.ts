import { randomUUID } from "crypto"
import * as uuid from 'uuid'

/**
 * random id for rooms
 *  */ 
export function generateRoomID() {
	

	return 'room-' + uuid.v4();
}