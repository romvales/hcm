import { test, expect } from 'bun:test'
import { Supabase_HCMWorkerService } from '.'
import { WorkerAddressType, WorkerGender } from '../../worker.d'

import {

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
import { User, createClient } from '@supabase/supabase-js'
import { Supabase_HCMWorkerOrganizationService } from './HCMWorkerService'
import { Database } from '../../database'

const localClient = createClient<Database>(
  process.env.SUPABASE_URL as string, 
  process.env.SUPABASE_LOCAL_SERVICE_ROLE_KEY as string,
)

const LOCALCLIENT_TEST_EMAIL = 'test@romvales.com'
const LOCALCLIENT_TEST_PASSWORD = 'testpassword'

export const createRandomWorker = () => {
  return {
    firstName: randFirstName(),
    lastName: randLastName(),
    middleName: randLastName(),
    username: randUserName(),
    email: randEmail(),
    gender: rand<WorkerGender>([ WorkerGender.MALE, WorkerGender.FEMALE, WorkerGender.OTHER ]),
  }
}

const setupHCMWorkerServiceTest = async () => {
  const workerService = new Supabase_HCMWorkerService(localClient)
  const _worker = createRandomWorker()
  const _creatorWorker = createRandomWorker()
  const worker = workerService.createWorker(_worker.email, _worker.username)
  const creatorWorker = workerService.createWorker(_creatorWorker.email, _creatorWorker.username)

  // Create mock user
  await localClient.auth.signUp({
    email: LOCALCLIENT_TEST_EMAIL,
    password: LOCALCLIENT_TEST_PASSWORD,
  })

  // Sign in the mock user
  const res = await localClient.auth.signInWithPassword({
    email: LOCALCLIENT_TEST_EMAIL,
    password: LOCALCLIENT_TEST_PASSWORD,
  })

  const user = res.data.user as User
  const id = user?.id as string

  return {
    workerService,
    _worker,
    worker,
    user,
    id,

    _creatorWorker,
    creatorWorker,
  }
}

const cleanHCMWorkerServiceTest = async (userId?: string) => {
  await localClient.auth.signOut()
  if (userId) await localClient.auth.admin.deleteUser(userId)
}

test ('> HCMWorkerSevice.createWorker()', async () => {
  const { workerService, _worker, id } = await setupHCMWorkerServiceTest()
  const worker = workerService.createWorker(_worker.email, _worker.username)

  expect(worker).not.toBeUndefined()
  expect(worker.email).toBe(_worker.email)
  expect(worker.username).toBe(_worker.username)

  await cleanHCMWorkerServiceTest(id)
})

test('> HCMWorkerService.addWorkerAddress()', async () => {
  const { workerService, worker, id } = await setupHCMWorkerServiceTest()

  expect(worker.addresses).toHaveLength(0)

  workerService
    .setTarget(worker)
    .addWorkerAddress({
      addrType: WorkerAddressType.HOME,
      city: randCity(),
      state: randState(),
      country: randCountry(),
      streetLines: [
        randStreetAddress(),
      ]
    })

  expect(worker.addresses).toHaveLength(1)
  await cleanHCMWorkerServiceTest(id)
})

test('> HCMWorkerService.getWorkerByUser()', async () => {
  const { 
    workerService, 
    worker, 
    user,
    _worker,
    id } = await setupHCMWorkerServiceTest()

  await workerService
    .setTarget(worker)
    .assignUserToWorker(user)
    .changeGender(_worker.gender)
    .changeName({
      firstName: _worker.firstName,
      lastName: _worker.lastName,
      middleName: _worker.middleName,
    })
    .changeEmailAddress(_worker.email)
    .changeUsername(_worker.username)
    .saveWorker()

  const getSavedWorker = await workerService.getWorkerByUser(user)

  expect(user).toBeDefined()
  expect(getSavedWorker).toBeDefined()
  expect(getSavedWorker?.firstName).toBe(_worker.firstName)
  expect(getSavedWorker?.lastName).toBe(_worker.lastName)
  expect(getSavedWorker?.email).toBe(_worker.email)
  expect(getSavedWorker?.userId).toBe(id)

  await cleanHCMWorkerServiceTest(id)
})

test('> HCMWorkerService.getWorkerById()', async () => {
  const {
    workerService, worker, _worker, id } = await setupHCMWorkerServiceTest()

  const savedWorker = await workerService
    .setTarget(worker)
    .changeName({
      firstName: _worker.firstName,
      lastName: _worker.lastName,
      middleName: _worker.middleName,
    })
    .changeGender(_worker.gender)
    .changeEmailAddress(_worker.email)
    .changeUsername(_worker.username)
    .saveWorker()

  const result = await workerService.getWorkerById(savedWorker.id)

  expect(result).toBeDefined()
  expect(result?.id).toBe(savedWorker.id as number)
  expect(result?.firstName).toBe(savedWorker.firstName as string)
  expect(result?.lastName).toBe(savedWorker.lastName as string)
  expect(result?.email).toBe(savedWorker.email as string)

  await localClient.from('workers').delete().match({ id: result?.id })
  await cleanHCMWorkerServiceTest(id)
})

test('> HCMWorkerService.deleteWorkerById()', async () => {
  const { 
    workerService, 
    worker, 
    _worker,
    id } = await setupHCMWorkerServiceTest()

  const savedWorker = await workerService
    .setTarget(worker)
    .changeName({
      firstName: _worker.firstName,
      lastName: _worker.lastName,
      middleName: _worker.middleName,
    })
    .changeGender(_worker.gender)
    .changeEmailAddress(_worker.email)
    .changeUsername(_worker.username)
    .saveWorker()

  const res = await workerService.deleteWorkerById(savedWorker.id)

  expect(await workerService.getWorkerById(savedWorker.id)).toBeUndefined()
  expect(res?.status).toBe(204)

  await cleanHCMWorkerServiceTest(id)
})

test('> HCMWorkerService.saveWorker()', async () => {
  const { workerService, worker, _worker, id } = await setupHCMWorkerServiceTest()

  const savedWorker = await workerService
    .setTarget(worker)
    .changeGender(_worker.gender)
    .changeName({
      firstName: _worker.firstName,
      lastName: _worker.lastName,
      middleName: _worker.middleName,
    })
    .changeEmailAddress(_worker.email)
    .changeUsername(_worker.username)
    .saveWorker()

  expect(savedWorker.email).toBe(worker.email)
  expect(savedWorker.firstName).toBe(_worker.firstName)
  expect(savedWorker.lastName).toBe(_worker.lastName)
  expect(savedWorker.middleName).toBe(_worker.middleName)
  expect(savedWorker.gender).toBe(_worker.gender)
  expect(savedWorker.username).toBe(_worker.username)
  expect(savedWorker.userId).toBeNull()
  
  await localClient.from('workers').delete().eq('email', _worker.email)
  await cleanHCMWorkerServiceTest(id)
})

test('> HCMWorkerService.saveWorker() # with session', async () => {
  const { 
    _creatorWorker,
    creatorWorker,
    user,
    id: creatorId,
    workerService,
    _worker,
    worker } = await setupHCMWorkerServiceTest()

  const creator = await workerService
    .setTarget(creatorWorker)
    .assignUserToWorker(user)
    .changeGender(_creatorWorker.gender)
    .changeName({
      firstName: _creatorWorker.firstName,
      lastName: _creatorWorker.lastName,
      middleName: _creatorWorker.middleName,
    })

    // Use the email of the authenticated test user
    .changeEmailAddress(user?.email as string)
    .changeUsername(_creatorWorker.username)
    .saveWorker()

  const createdWorker = await workerService
    .setTarget(worker)
    .changeGender(_worker.gender)
    .changeName({
      firstName: _worker.firstName,
      lastName: _worker.lastName,
      middleName: _worker.middleName,
    })
    .changeEmailAddress(_worker.email)
    .changeUsername(_worker.username)
    .saveWorker()

  expect(createdWorker.createdById == creator.id).toBeTrue()
  expect(createdWorker.userId).toBeNull()
  expect(creator.userId).toBe(creatorId)
  expect(user?.email).toBe(creator.email)
  expect(creator.firstName).toBe(_creatorWorker.firstName)
  expect(creator.lastName).toBe(_creatorWorker.lastName)
  expect(createdWorker.firstName).toBe(_worker.firstName)
  expect(createdWorker.lastName).toBe(_worker.lastName)
  expect(createdWorker.email).toBe(_worker.email)

  await localClient.from('workers').delete().eq('email', createdWorker.email)
  await cleanHCMWorkerServiceTest(creatorId)
})

test('> HCMWorkerService.saveWorkerIdentityCard()', async () => {
  const ID_NAME = 'Personal Driver\'s License (1)'
  const { workerService, _worker, worker, id } = await setupHCMWorkerServiceTest()

  const createdWorker = await workerService
    .setTarget(worker)
    .changeGender(_worker.gender)
    .changeName({
      firstName: _worker.firstName,
      lastName: _worker.lastName,
      middleName: _worker.middleName,
    })
    .changeEmailAddress(_worker.email)
    .changeUsername(_worker.username)
    .saveWorker()

  const idCard = await workerService.saveWorkerIdentityCard({ 
    workerId: createdWorker.id,
    name: ID_NAME,
  })

  expect(idCard).toBeDefined()
  expect(idCard.name).toBe(ID_NAME)

  await localClient.from('workersIdentityCards').delete().eq('name', ID_NAME)
  await workerService.deleteWorkerById(createdWorker.id)
  await cleanHCMWorkerServiceTest(id)
})

test('> HCMWorkerService.getIdentityCards()', async () => {


})

test('> HCMWorkerService.deleteIdentityCardById()', async () => {


})

test ('> HCMWorkerOrganizationService.getOrganizations()', async () => {
  const workerOrganizationService = new Supabase_HCMWorkerOrganizationService(localClient)
  const workerService = new Supabase_HCMWorkerService(localClient)
  const _worker = createRandomWorker()
  const worker = workerService.createWorker(_worker.email, _worker.username)

  const createdWorker = await workerService
    .setTarget(worker)
    .changeGender(_worker.gender)
    .changeName({
      firstName: _worker.firstName,
      lastName: _worker.lastName,
      middleName: _worker.middleName,
    })
    .changeEmailAddress(_worker.email)
    .changeUsername(_worker.username)
    .saveWorker()

  workerOrganizationService
    .setTarget(createdWorker)
    .getOrganizations()
  

  await localClient.from('workers').delete().match({ id: createdWorker.id })
})