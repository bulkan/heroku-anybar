const { exec } = require('child_process');
const ora = require('ora');
const anybar = require('anybar');
const spinner = ora('wat').start();
spinner.color = 'yellowBright';

process.stdin.resume();

function exitHandler(err) {
  if (err) {
    console.error(err);
  }
  anybar('exclamation')
    .then(process.exit);
}

[ 'exit', 'SIGINT', 'SIGUSR1', 'SIGUSR2', 'uncaughtException'].forEach(sig => process.on(sig, exitHandler));

function updateStatus() {
  spinner.text = 'updating';
  anybar('blue');

  const PIPELINE = process.env.PIPELINE;

  exec(`heroku ci --json --pipeline ${PIPELINE}`, (err, stdout, stderr) => {
    if (err) {
      anybar('exclamation');
      spinner.color = 'red';
      spinner.text = err;
      console.error(stdout);
      return;
    }

    const build = JSON.parse(stdout)[0];
    spinner.text = `${build.commit_sha} - ${build.status}`;
    spinner.color = 'yellowBright';

    switch (build.status) {
      case 'building':
        anybar('orange');
        spinner.color = 'cyan';
        spinner.text = build.status;
        break;
      case 'running':
        anybar('yellow');
        spinner.text = build.status;
        spinner.color = 'yellowBright';
        break;
      case 'creating':
        anybar('orange');
        spinner.text = build.status;
        spinner.color = 'yellowBright';
        break;
      case 'pending':
        anybar('orange');
        spinner.text = build.status;
        spinner.color = 'yellowBright';
        break;
      case 'failed':
        anybar('red');
        spinner.text = build.status;
        spinner.color = 'redBright';
        break;
      case 'errored':
        anybar('red');
        spinner.text = build.status;
        spinner.color = 'redBright';
        break;
      case 'succeeded':
        anybar('green');
        spinner.color = 'green';
        break;
      default:
        anybar('exclamation');
        spinner.text = `no idea ${build.status}`
        spinner.color = 'redBright';
        break;
    }
  });
}

if (require.main === module) {
  updateStatus();
  setInterval(updateStatus, 60000);
}

