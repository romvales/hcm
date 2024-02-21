import { test, expect, afterEach } from 'bun:test'
import { Supabase_HCMWorkerService } from '.'
import { Worker, WorkerAddressType, WorkerGender } from '../../../src'

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
import { Database } from '../../../src'

const localClient = createClient<Database>(
  process.env.SUPABASE_URL as string, 
  process.env.SUPABASE_LOCAL_SERVICE_ROLE_KEY as string,
)

const LOCALCLIENT_TEST_EMAIL = 'test@romvales.com'
const LOCALCLIENT_TEST_PASSWORD = 'testpassword'

let userId: string

afterEach(async () => {
  await cleanHCMWorkerServiceTest(userId)
})

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

  if (user) userId = user?.id

  return {
    workerService,
    _worker,
    worker,
    user,

    _creatorWorker,
    creatorWorker,
  }
}

export const setupFakeWorkerFields = (
  workerService: Supabase_HCMWorkerService, 
  target: Worker,
  user?: User | null, 
  workerFields?: any) => {

  workerService
    .setTarget(target)
    .changeGender(workerFields.gender)
    .changeName({
      firstName: workerFields.firstName,
      lastName: workerFields.lastName,
      middleName: workerFields.middleName,
    })
    .changeEmailAddress(workerFields.email)
    .changeUsername(workerFields.username)

  if (user) {
    workerService.assignUserToWorker(user)
  }

  if (user?.email) {
    workerService.changeEmailAddress(user.email)
    workerFields.email = user.email
  }

  return workerService
}

const expectCommonFieldsToBeDefined = (worker?: Worker, _worker?: any) => {
  if (!worker) return

  expect(worker.email).toBe(_worker.email)
  expect(worker.username).toBe(_worker.username)
}

const cleanHCMWorkerServiceTest = async (userId?: string) => {
  await localClient.auth.signOut()
  if (userId) await localClient.auth.admin.deleteUser(userId)
}

test ('> HCMWorkerSevice.createWorker()', async () => {
  const { workerService, _worker } = await setupHCMWorkerServiceTest()
  const worker = workerService.createWorker(_worker.email, _worker.username)

  expect(worker).not.toBeUndefined()
  expectCommonFieldsToBeDefined(worker, _worker)
})

test('> HCMWorkerService.addWorkerAddress()', async () => {
  const { workerService, worker } = await setupHCMWorkerServiceTest()

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
})

test('> HCMWorkerService.getWorkerByUser()', async () => {
  const { 
    workerService, 
    worker, 
    user,
    _worker } = await setupHCMWorkerServiceTest()

  await setupFakeWorkerFields(workerService, worker, user, _worker).saveWorker()

  const getSavedWorker = await workerService.getWorkerByUser(user)

  expect(user).toBeDefined()
  expect(getSavedWorker).toBeDefined()
  expectCommonFieldsToBeDefined(getSavedWorker, _worker)
  expect(getSavedWorker?.firstName).toBe(_worker.firstName)
  expect(getSavedWorker?.lastName).toBe(_worker.lastName)
  expect(getSavedWorker?.userId).toBe(userId)
})

test('> HCMWorkerService.getWorkerById()', async () => {
  const {
    workerService, worker, _worker } = await setupHCMWorkerServiceTest()

  const savedWorker = await setupFakeWorkerFields(workerService, worker, undefined, _worker)
    .saveWorker()

  const result = await workerService.getWorkerById(savedWorker.id)

  expect(result).toBeDefined()
  expect(result?.id).toBe(savedWorker.id as number)
  expectCommonFieldsToBeDefined(result, savedWorker)
  expect(result?.firstName).toBe(savedWorker.firstName as string)
  expect(result?.lastName).toBe(savedWorker.lastName as string)

  await localClient.from('workers').delete().match({ id: result?.id })
})

test('> HCMWorkerService.deleteWorkerById()', async () => {
  const { 
    workerService, 
    worker, 
    _worker } = await setupHCMWorkerServiceTest()

  const savedWorker = await setupFakeWorkerFields(workerService, worker, undefined, _worker)
    .saveWorker()

  const res = await workerService.deleteWorkerById(savedWorker.id)

  expect(await workerService.getWorkerById(savedWorker.id)).toBeUndefined()
  expect(res?.status).toBe(204)
})

test('> HCMWorkerService.saveWorker()', async () => {
  const { workerService, worker, _worker } = await setupHCMWorkerServiceTest()

  const savedWorker = await setupFakeWorkerFields(workerService, worker, undefined, _worker)
    .saveWorker()

  expectCommonFieldsToBeDefined(savedWorker, _worker)
  expect(savedWorker.firstName).toBe(_worker.firstName)
  expect(savedWorker.lastName).toBe(_worker.lastName)
  expect(savedWorker.middleName).toBe(_worker.middleName)
  expect(savedWorker.gender).toBe(_worker.gender)
  expect(savedWorker.username).toBe(_worker.username)
  expect(savedWorker.userId).toBeNull()
  
  await localClient.from('workers').delete().eq('email', _worker.email)
})

test('> HCMWorkerService.saveWorker() # with session', async () => {
  const { 
    _creatorWorker,
    creatorWorker,
    user,
    workerService,
    _worker,
    worker } = await setupHCMWorkerServiceTest()

  const creator = await setupFakeWorkerFields(workerService, creatorWorker, user, _creatorWorker)
    .saveWorker()

  const createdWorker = await setupFakeWorkerFields(workerService, worker, undefined, _worker)
    .saveWorker()

  expect(createdWorker.createdById == creator.id).toBeTrue()
  expect(createdWorker.userId).toBeNull()
  expect(creator.userId).toBe(userId)
  expect<string | undefined>(creator.email).toBe(user.email)
  expectCommonFieldsToBeDefined(creator, _creatorWorker)
  expect(creator.firstName).toBe(_creatorWorker.firstName)
  expect(creator.lastName).toBe(_creatorWorker.lastName)
  expectCommonFieldsToBeDefined(createdWorker, _worker)
  expect(createdWorker.firstName).toBe(_worker.firstName)
  expect(createdWorker.lastName).toBe(_worker.lastName)

  await localClient.from('workers').delete().eq('email', createdWorker.email)
})

test('> HCMWorkerService.saveWorkerIdentityCard()', async () => {
  const ID_NAME = 'Personal Driver\'s License (1)'
  const { workerService, _worker, worker } = await setupHCMWorkerServiceTest()

  const createdWorker = 
    await setupFakeWorkerFields(workerService, worker, undefined, _worker)
      .saveWorker()

  const idCard = await workerService.saveWorkerIdentityCard({ 
    workerId: createdWorker.id,
    name: ID_NAME,
  })

  expect(idCard).toBeDefined()
  expect(idCard.name).toBe(ID_NAME)

  await localClient.from('workersIdentityCards').delete().eq('name', ID_NAME)
  await workerService.deleteWorkerById(createdWorker.id)
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