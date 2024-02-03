import { Organization } from '../../index.d'
import { 
  HCMWorkerService, 
  Role, 
  Team, 
  Worker, 
  WorkerAddress, 
  WorkerGender,
  WorkerIdentityCard,
  WorkerPayCycle,
  WorkerStatus,
  WorkerType } from '../../worker.d'

import { SupabaseClient } from '@supabase/supabase-js'

export class Supabase_HCMWorkerService extends HCMWorkerService {

  constructor(
    private client?: SupabaseClient
  ) {
    super()
    
  }

  // createWorker() -> Creates a new worker record without saving it to the database
  createWorker(params: {
    pictureUrl?: string,
    createdById: number,
    firstName: string, 
    middleName?: string, 
    lastName: string,
    email: string,
    username: string,
    gender: WorkerGender,
    mobileNumber?: string,
    birthdate?: number,
    addresses?: WorkerAddress[],
  }): Worker {
    const { 
      createdById, 
      firstName, 
      lastName, 
      middleName, 
      email, 
      username, 
      mobileNumber,
      gender,
      birthdate,
      addresses } = params

    const out: Worker = {
      createdById,
      firstName,
      lastName,
      middleName,
      email,
      username,
      gender,
      mobile: mobileNumber,
      birthdate,
      addresses: addresses ?? [],
    }
    
    return out
  }

  async getWorkerById(id: number): Promise<Worker | void> {
    
  }

  async deleteWorkerById<T>(id: number): Promise<T> {
    
  }

  async saveWorker(updator: Worker | null, worker: Worker) {
    const supabaseClient = this.client

    if (!supabaseClient) {
      return
    }

    // When the updator is the same as the worker, don't update
    // who updated the record.
    if (updator && updator.id !== worker.id) {    
      worker.updatedById = updator.id
    }

    

  } 

  // Adds new identity cards for a worker
  addIdentityCards(worker: Worker, identityCards: WorkerIdentityCard[]) {
    if (!worker.identityCards?.length) worker.identityCards = []
    worker.identityCards.push(...identityCards)
    return this
  }

  // Pushes a new address for a worker
  addWorkerAddress(worker: Worker, address: WorkerAddress) {
    if (!worker.addresses?.length) worker.addresses = []
    worker.addresses?.unshift(address)
    return this
  }

}