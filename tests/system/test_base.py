from errlogbeat import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Errlogbeat normally
        """
        self.render_config_template(
            path=os.path.abspath(self.working_dir) + "/log/*"
        )

        errlogbeat_proc = self.start_beat()
        self.wait_until(lambda: self.log_contains("errlogbeat is running"))
        exit_code = errlogbeat_proc.kill_and_wait()
        assert exit_code == 0
