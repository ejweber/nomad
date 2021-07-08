import { currentURL } from '@ember/test-helpers';
import { module, test } from 'qunit';
import { setupApplicationTest } from 'ember-qunit';
import { setupMirage } from 'ember-cli-mirage/test-support';
import a11yAudit from 'nomad-ui/tests/helpers/a11y-audit';
import ServerDetail from 'nomad-ui/tests/pages/servers/detail';
import formatHost from 'nomad-ui/utils/format-host';

let agent;

module('Acceptance | server detail', function(hooks) {
  setupApplicationTest(hooks);
  setupMirage(hooks);

  hooks.beforeEach(async function() {
    server.createList('agent', 3);
    agent = server.db.agents[0];
    await ServerDetail.visit({ name: agent.member.Name });
  });

  test('it passes an accessibility audit', async function(assert) {
    await a11yAudit(assert);
  });

  test('visiting /servers/:server_name', async function(assert) {
    console.log('agent:  ', agent);
    assert.equal(currentURL(), `/servers/${encodeURIComponent(agent.member.Name)}`);
    assert.equal(document.title, `Server ${agent.member.Name} - Nomad`);
  });

  test('when the server is the leader, the title shows a leader badge', async function(assert) {
    assert.ok(ServerDetail.title.includes(agent.member.Name));
    assert.ok(ServerDetail.hasLeaderBadge);
  });

  test('the details ribbon displays basic information about the server', async function(assert) {
    assert.ok(ServerDetail.serverStatus.includes(agent.member.Status));
    assert.ok(
      ServerDetail.address.includes(formatHost(agent.member.Address, agent.member.Tags.port))
    );
    assert.ok(ServerDetail.datacenter.includes(agent.member.Tags.dc));
  });

  test('the server detail page should list all tags for the server', async function(assert) {
    const tags = Object.keys(agent.member.Tags)
      .map(name => ({ name, value: agent.member.Tags[name] }))
      .sortBy('name');

    assert.equal(ServerDetail.tags.length, tags.length, '# of tags');
    ServerDetail.tags.forEach((tagRow, index) => {
      const tag = tags[index];
      assert.equal(tagRow.name, tag.name, `Label: ${tag.name}`);
      assert.equal(tagRow.value, tag.value, `Value: ${tag.value}`);
    });
  });

  test('when the server is not the leader, there is no leader badge', async function(assert) {
    await ServerDetail.visit({ name: server.db.agents[1].member.Name });
    assert.notOk(ServerDetail.hasLeaderBadge);
  });

  test('when the server is not found, an error message is shown, but the URL persists', async function(assert) {
    await ServerDetail.visit({ name: 'not-a-real-server' });

    assert.equal(currentURL(), '/servers/not-a-real-server', 'The URL persists');
    assert.equal(ServerDetail.error.title, 'Not Found', 'Error message is for 404');
  });
});
