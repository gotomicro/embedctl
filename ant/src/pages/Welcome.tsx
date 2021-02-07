import React from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { Card, Alert, Typography } from 'antd';

export default (): React.ReactNode => {
  return (
    <PageContainer>
      <Card>
        <Alert
          message="命令行嵌入网站。"
          type="success"
          showIcon
          banner
          style={{
            margin: -12,
            marginBottom: 24,
          }}
        />
        <Typography.Text strong>
          <a
            href="https://github.com/gotomicro/egoctl"
            rel="noopener noreferrer"
            target="__blank"
          >
            欢迎使用
          </a>

          Egoctl
        </Typography.Text>
      </Card>
    </PageContainer>
  );
};
