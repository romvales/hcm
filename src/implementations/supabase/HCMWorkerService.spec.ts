import { test, expect } from 'bun:test'
import { Supabase_HCMWorkerService } from '.'
import { WorkerAddressType, WorkerGender } from '../../worker.d'

import {

  randNumber,
  randEmail,
  randUserName,
  randFirstName,
  randLastName,
  rand,

  randState,
  randCity,
  randCountry,
  randStreetAddress,

} from '@ngneat/falso'
import { createClient } from '@supabase/supabase-js'

const localClient = createClient(
  process.env.SUPABASE_URL as string, 
  process.env.SUPABASE_LOCAL_ANON_KEY as string,
)

function createRandomWorker() {
  return {
    createdById: randNumber(),
    firstName: randFirstName(),
    lastName: randLastName(),
    middleName: randLastName(),
    username: randUserName(),
    email: randEmail(),
    gender: rand<WorkerGender>([ WorkerGender.MALE, WorkerGender.FEMALE, WorkerGender.OTHER ]),
  }
}

test ('> HCMWorkerSevice.createWorker()', () => {
  const workerService = new Supabase_HCMWorkerService()
  
  const worker = workerService.createWorker(createRandomWorker())

  expect(worker).not.toBeUndefined()
})

test('> HCMWorkerService.addWorkerAddress()', () => {
  const workerService = new Supabase_HCMWorkerService()

  const worker = workerService.createWorker(createRandomWorker())

  expect(worker.addresses).toHaveLength(0)

  workerService
    .addWorkerAddress(worker, {
      addrType: WorkerAddressType.HOME,
      city: randCity(),
      state: randState(),
      country: randCountry(),
      streetLines: [
        randStreetAddress(),
      ]
    })
    

  expect(worker.addresses).toHaveLength(1)
})

test('> HCMWorkerService.saveWorker()', async () => {
  const workerService = new Supabase_HCMWorkerService(localClient)
  const worker = workerService.createWorker(createRandomWorker())

  workerService.saveWorker(null, worker)
})